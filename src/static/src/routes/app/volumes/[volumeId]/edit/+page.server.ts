import { makeApiUrl } from '$lib/proxy';
import type { Volume } from '$lib/types';
import { collectFormValues } from '$lib/util';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad, Actions } from './$types';

export const load = (async ({ fetch, params }) => {
	const json = (url: string) => fetch(makeApiUrl(url)).then((res) => res.json());

	return {
		volume: await json(`/api/volumes/${params.volumeId}`),
	} as {
		volume: Volume;
	};
}) satisfies PageServerLoad;

export const actions: Actions = {
	update: async ({ request, fetch, params }) => {
		const data = collectFormValues(await request.formData());

		data.currentPage = parseInt(data.currentPage, 10);
		data.pagesRead = parseInt(data.pagesRead, 10);

		await fetch(makeApiUrl(`/api/volumes/${params.volumeId}`), {
			method: 'PATCH',
			body: JSON.stringify(data),
			headers: {
				'Content-Type': 'application/json',
			},
		});

		throw redirect(301, `/app/volumes/${params.volumeId}`);
	},
};
