export function apiPath(path: string) {
	//in prod mode the go api hosts the static files, in dev
	//vite hosts them on :3000 so we need to specify the port
	const host = import.meta.env.PROD ? '' : 'http://localhost:5004';
	return `${host}${path}`;
}

export async function apiRequest<T>(url: string, method = 'GET', body?: any) {
	const options = ['POST', 'PATCH'].includes(method)
			? {
					headers: { 'Content-Type': 'application/json' },
					body: JSON.stringify(body),
			  }
			: {},
		res = await fetch(apiPath(`/api${url}`), {
			method,
			...options,
			credentials: 'include',
		});

	return {
		status: res.status,
		ok: res.ok,
		body: (await res.json().catch(() => null)) as T,
	};
}
