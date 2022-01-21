export async function apiRequest(url: string, method = 'GET', body?: any) {
	const options = ['POST', 'PATCH'].includes(method)
			? {
					headers: { 'Content-Type': 'application/json' },
					body: JSON.stringify(body),
			  }
			: {},
		//in prod mode the go api hosts the static files, in dev
		//vite hosts them on :3000.
		host = import.meta.env.PROD ? '' : 'http://localhost:4004';

	return (
		await fetch(`${host}/api${url}`, {
			method,
			...options,
		})
	).json();
}
