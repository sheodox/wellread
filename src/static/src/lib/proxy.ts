import fetch from 'node-fetch';

const API_HOST = 'localhost:5004';

export interface ProxiedResponse<T> {
	body: T;
	status: number;
	headers: Record<string, string>;
}

export const makeApiUrl = (pathname: string) => {
	return `http://${API_HOST}${pathname}`;
};

export const proxy = async <T>(request: Request, method: 'GET' | 'POST' = 'GET'): Promise<ProxiedResponse<T>> => {
	const url = new URL(request.url);
	url.host = API_HOST;

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
