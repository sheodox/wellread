import { XIcon } from '@heroicons/react/outline';
import { Empty } from './Empty';
import { Spinner } from './Spinner';
import { ReadingHistory as ReadingHistoryType, useStore } from './state/data';

export function ReadingHistory() {
	const { readingHistory, readingHistoryLoading, deleteReadingHistory } = useStore(),
		deleteHistory = async (history: ReadingHistoryType) => {
			if (confirm(`Are you sure you want to delete the page ${history.currentPage} reading history?`)) {
				deleteReadingHistory(history.id);
			}
		};

	return (
		<div className="md:w-48">
			<h1 className="pb-2 mb-4 border-b border-slate-700 flex items-baseline">
				Reading History
				<div className="flex justify-center ml-4">
					<Spinner show={readingHistoryLoading} />
				</div>
			</h1>
			{!readingHistoryLoading && !readingHistory.length && <Empty />}
			{!readingHistoryLoading && (
				<ul className="flex-row md:flex-col">
					{readingHistory.map((h, i) => {
						const increase = i < readingHistory.length - 1 ? h.currentPage - readingHistory[i + 1].currentPage : null;

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
								<div className="flex justify-between items-end">
									<span className="text-slate-400">{new Date(h.createdAt).toLocaleDateString()}</span>
									<button
										className="ml-4 opacity-40 hover:opacity-100 hover:text-red-400 transition-all"
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
