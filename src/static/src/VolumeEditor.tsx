import { useAtom } from 'jotai';
import { ChangeEvent, useEffect, useState } from 'react';
import { apiRequest } from './state/api';
import { useSelectedSeries, useSelectedVolume, volumesAtom } from './state/data';

export function VolumeEditor() {
	const volume = useSelectedVolume(),
		setVolumes = useAtom(volumesAtom)[1],
		series = useSelectedSeries(),
		[notes, setNotes] = useState(volume?.notes),
		[currentPage, setCurrentPage] = useState(volume?.currentPage),
		[pageError, setPageError] = useState(false),
		[saving, setSaving] = useState(false);

	function resetStateToVolume() {
		if (volume) {
			setNotes(volume.notes);
			setCurrentPage(volume.currentPage);
		}
	}

	useEffect(() => {
		resetStateToVolume();
	}, [volume]);

	if (!volume || !series) {
		return null;
	}

	const save = async () => {
			setSaving(true);
			const volumes = await apiRequest(`/series/${series.id}/volumes/${volume.id}`, 'PATCH', {
				name: volume.name,
				notes,
				currentPage,
			});
			setVolumes(volumes);

			setSaving(false);
		},
		onPageChange = (e: ChangeEvent<HTMLInputElement>) => {
			const newPages = e.target.value;

			if (/^\d+$/.test(newPages)) {
				setCurrentPage(+newPages);
				setPageError(false);
			} else {
				setPageError(true);
			}
		},
		inputClasses =
			'rounded-md border border-zinc-700 bg-transparent bg-zinc-800 focus:outline-none focus:border-sky-500 p-2',
		saveable = !pageError && (currentPage !== volume.currentPage || notes !== volume.notes);

	return (
		<div className="flex-1 flex flex-col p-6">
			<div className="flex justify-between items-center">
				<h1 className="text-4xl">
					{volume.name} - {series.name}
				</h1>
				<button
					onClick={save}
					className={`rounded p-2 font-bold transition-colors border-2 ${
						saveable ? 'border-sky-400 text-sky-400' : 'border-transparent text-zinc-600'
					}`}
					disabled={saving || !saveable}
				>
					Save
				</button>
			</div>
			<label htmlFor="volume-current-page" className="mt-4">
				Current Page
			</label>
			<input
				id="volume-current-page"
				defaultValue={volume.currentPage}
				className={inputClasses}
				onChange={onPageChange}
			/>
			{pageError && <small className="text-red-400">Must enter a valid page number!</small>}
			<label htmlFor="volume-notes" className="mt-4">
				Notes
			</label>
			<textarea
				id="volume-notes"
				defaultValue={volume.notes}
				className={inputClasses + ' flex-1 resize-none'}
				onChange={(event) => setNotes(event.target.value)}
			></textarea>
		</div>
	);
}
