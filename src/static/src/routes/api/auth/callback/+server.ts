import { proxy } from '$lib/proxy';
import type { RequestHandler } from './$types';

export const POST = (async ({ request }) => {
	const p = await proxy<unknown>(request, 'POST');

	return new Response(null, { status: p.status, headers: p.headers });
}) satisfies RequestHandler;
