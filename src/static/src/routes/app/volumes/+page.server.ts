import { makeApiUrl } from '$lib/proxy';
import type { Volume, PagedResponse } from '$lib/types';
import type { PageServerLoad } from './$types';

export const load = (async ({ fetch, url }) => {
	const pageNum = url.searchParams.get('page') ?? '1',
		status = url.searchParams.get('status') ?? '',
		name = url.searchParams.get('name') ?? '';

	return {
		volumes: {
			...(await (
				await fetch(makeApiUrl(`/api/volumes?page=${pageNum}&status=${status}&name=${encodeURIComponent(name)}`))
			).json()),
			filter: {
				status,
				name,
			},
		},
	} as {
		volumes: PagedResponse<Volume[]>;
	};
}) satisfies PageServerLoad;
