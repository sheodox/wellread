import { ChangeEvent, ForwardedRef, forwardRef, useEffect, useRef, useState } from 'react';
import { useSelectedSeries, useSelectedVolume, useStore } from './state/data';
import StatusBadge from './StatusBadge';
import { theme } from './theme';

const readingStatuses = ['planning', 'reading', 'completed', 'dropped'];

function capitalize(str: string) {
	return str.charAt(0).toUpperCase() + str.substring(1);
}

export function VolumeEditor() {
	const { updateVolume, loadVolumes } = useStore(),
		selectedSeries = useSelectedSeries(),
		selectedVolume = useSelectedVolume(),
		[editing, setEditing] = useState(false),
		[notes, setNotes] = useState(selectedVolume?.notes),
		[currentPage, setCurrentPage] = useState(selectedVolume?.currentPage),
		[pageError, setPageError] = useState(false),
		[pagesReadError, setPagesReadError] = useState(false),
		pagesReadRef = useRef<HTMLInputElement>(null),
		[saving, setSaving] = useState(false),
		[status, setStatus] = useState('planning'),
		[pagesRead, setPagesRead] = useState(0),
		resetState = (force = false) => {
			if ((force || !editing) && selectedVolume) {
				setNotes(selectedVolume.notes);
				setCurrentPage(selectedVolume.currentPage);
				setStatus(selectedVolume.status);
				setPagesReadError(false);
			}
		};

	useEffect(resetState, [selectedVolume]);

	if (!selectedVolume || !selectedSeries) {
		return null;
	}

	const save = async () => {
			if (typeof notes === 'string' && typeof currentPage === 'number') {
				setSaving(true);
				await updateVolume(selectedVolume.id, { notes, name: selectedVolume.name, currentPage, status, pagesRead });
				setSaving(false);
				setEditing(false);
				setPagesRead(0);
			}
		},
		startEditing = async () => {
			//ensure we're up to date before resuming any edit
			await loadVolumes(selectedSeries.id);
			setEditing(true);
		},
		cancelEditing = () => {
			resetState(true);
			setEditing(false);
		},
		onPageChange = (newPages: number) => {
			const pages = +newPages;
			setCurrentPage(pages);

			const pageDelta = Math.max(0, pages - selectedVolume.currentPage);
			setPagesRead(pageDelta);
			if (pagesReadRef.current) {
				pagesReadRef.current.value = '' + pageDelta;
			}
		},
		inputClasses = 'rounded-md border border-slate-600 bg-slate-700 focus:outline-none focus:border-sky-500 p-2',
		editingButtons = editing ? (
			<div>
				<button className={theme.button.primary} onClick={save} disabled={pageError || pagesReadError || saving}>
					Save
				</button>
				<button className={`ml-1 ${theme.button.secondary}`} onClick={cancelEditing}>
					Cancel
				</button>
			</div>
		) : (
			<div>
				<button className={theme.button.primary} onClick={startEditing}>
					Edit
				</button>
			</div>
		);

	return (
		<div className="flex-1 flex flex-col md:min-w-[40rem]">
			<div className="mb-6 flex items-center">
				<h1 className="text-4xl mr-4">
					{selectedVolume.name} - {selectedSeries.name}
				</h1>
				<StatusBadge status={status} size="large" />
			</div>
			<div className="flex justify-between items-start">
				<div className="mb-4">
					<p>Current Page</p>
					<p className="text-4xl">{selectedVolume.currentPage}</p>
				</div>
				{editingButtons}
			</div>
			{editing ? (
				<>
					<div className="grid grid-cols-2 md:grid-cols-3 gap-4 items-end">
						<div className="flex flex-col">
							<label htmlFor="volume-current-page">New current page</label>
							<NumberInput
								id="volume-current-page"
								defaultValue={selectedVolume.currentPage}
								className={`${inputClasses} w-full`}
								onChange={onPageChange}
								onErrorStatusChange={setPageError}
							/>
						</div>
						<div className="flex flex-col">
							<label htmlFor="volume-current-page">Newly read pages</label>
							<NumberInput
								id="volume-current-page"
								defaultValue={pagesRead}
								className={`${inputClasses} w-full`}
								onChange={(pages) => setPagesRead(pages)}
								onErrorStatusChange={setPagesReadError}
								ref={pagesReadRef}
							/>
						</div>
						<div className="flex flex-col">
							<label htmlFor="volume-status">Status</label>
							<select
								id="volume-status"
								defaultValue={status}
								className={`${inputClasses} w-full`}
								onChange={(e) => setStatus(e.target.value)}
							>
								{readingStatuses.map((status) => {
									return (
										<option key={status} value={status}>
											{capitalize(status)}
										</option>
									);
								})}
							</select>
						</div>
					</div>

					<div className="mt-4 flex flex-1 flex-col">
						<div className="flex justify-between items-end mb-1">
							<label htmlFor="volume-notes">Notes</label>
						</div>
						<textarea
							id="volume-notes"
							defaultValue={selectedVolume.notes}
							className={inputClasses + ' flex-1 resize-none'}
							onChange={(event) => setNotes(event.target.value)}
						></textarea>
					</div>
				</>
			) : (
				<>
					<p className="mb-1 border-b border-slate-700">Notes</p>
					<p className="whitespace-pre-line">
						{notes ? notes : <span className="text-slate-400 italic">You haven't written any notes yet.</span>}
					</p>
				</>
			)}
		</div>
	);
}

interface NumberInputProps {
	onChange: (num: number) => void;
	onErrorStatusChange: (valid: boolean) => void;
	defaultValue: any;
	className: string;
	id: string;
}

const NumberInput = forwardRef((props: NumberInputProps, ref: ForwardedRef<HTMLInputElement>) => {
	const [invalidNumber, setInvalidNumber] = useState(false),
		onChange = (e: ChangeEvent<HTMLInputElement>) => {
			const newValue = e.target.value;

			if (/^\d+$/.test(newValue)) {
				props.onChange(+newValue);
				props.onErrorStatusChange(false);
				setInvalidNumber(false);
			} else {
				props.onErrorStatusChange(true);
				setInvalidNumber(true);
			}
		};
	return (
		<>
			<input
				onChange={onChange}
				defaultValue={props.defaultValue}
				className={props.className}
				id={props.id}
				type="number"
				ref={ref}
			/>
			{invalidNumber && <small className="text-red-400">Must enter a valid number!</small>}
		</>
	);
});
