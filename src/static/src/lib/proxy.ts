import fetch from 'node-fetch';
import { dev } from '$app/environment';

const API_HOST = !dev ? 'api:5004' : 'localhost:5004';

export interface ProxiedResponse<T> {
	body: T;
	status: number;
	headers: Record<string, string>;
}

export const makeApiUrl = (pathname: string) => {
	const host = !dev ? 'http://api:5004' : 'http://localhost:5004';
	return `http://${API_HOST}${pathname}`;
};

export const proxy = async <T>(request: Request, method: 'GET' | 'POST' = 'GET'): Promise<ProxiedResponse<T>> => {
	const url = new URL(request.url);
	url.host = API_HOST;
	// inside the docker network the api server uses http
	url.protocol = 'http:';

	const headers = {
		cookie: request.headers.get('cookie') ?? '',
	};

	const proxied = await fetch(url.toString(), {
		method,
		...(method === 'POST'
			? {
					headers: {
						...headers,
						'Content-Type': request.headers.get('Content-Type') ?? 'text/plain',
					},
					body: await request.text(),
			  }
			: {
					headers,
			  }),
	});

	const responseBody = (await proxied.json()) as T;

	return {
		body: responseBody,
		status: proxied.status,
		headers: {
			'Set-Cookie': proxied.headers.get('set-cookie') ?? '',
		},
	};
};
