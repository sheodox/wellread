import { XIcon } from '@heroicons/react/outline';
import { useAtom } from 'jotai';
import { useParams } from 'react-router-dom';
import { Empty } from './Empty';
import { Spinner } from './Spinner';
import { apiRequest } from './state/api';
import { ReadingHistory as ReadingHistoryType, readingHistoryAtom, readingHistoryLoadingAtom } from './state/data';

export function ReadingHistory() {
	const { seriesId, volumeId } = useParams(),
		[history, setHistory] = useAtom(readingHistoryAtom),
		[historyLoading, setHistoryLoading] = useAtom(readingHistoryLoadingAtom),
		deleteHistory = async (history: ReadingHistoryType) => {
			if (confirm(`Are you sure you want to delete the page ${history.currentPage} reading history?`)) {
				setHistoryLoading(true);
				const h = await apiRequest(`/series/${seriesId}/volumes/${volumeId}/history/${history.id}`, 'DELETE');
				setHistory(h);
				setHistoryLoading(false);
			}
		};

	return (
		<div className="mx-9">
			<h1 className="pb-2 mb-4 border-b border-zinc-700">Page History</h1>
			{historyLoading && (
				<div className="flex justify-center mt-4">
					<Spinner />
				</div>
			)}
			{!historyLoading && !history.length && <Empty />}
			{!historyLoading && (
				<ul>
					{history.map((h, i) => {
						const increase = i < history.length - 1 ? h.currentPage - history[i + 1].currentPage : null;

						return (
							<li key={h.id} className="mb-7">
								<div className="flex justify-between items-baseline">
									<span className="text-2xl">{h.currentPage}</span>{' '}
									{increase && (
										<>
											{increase > 0 ? (
												<span className="text-green-400">+{increase}</span>
											) : (
												<span className="text-red-400">{increase}</span>
											)}
										</>
									)}
								</div>
								<div className="flex items-end">
									<span className="text-zinc-400">{new Date(h.createdAt).toLocaleDateString()}</span>
									<button
										className="ml-4 opacity-10 hover:opacity-100 hover:text-red-400 transition-all"
										onClick={() => deleteHistory(h)}
										title="Delete this reading history"
									>
										<XIcon className="h-5" />
										<span className="sr-only"> Delete </span>
									</button>
								</div>
							</li>
						);
					})}
				</ul>
			)}
		</div>
	);
}
