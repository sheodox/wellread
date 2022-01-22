import { useAtom } from 'jotai';
import { Fragment, useEffect } from 'react';
import { BrowserRouter, Route, Routes, useParams } from 'react-router-dom';
import { SeriesSelector, VolumeSelector } from './Selector';
import { apiRequest } from './state/api';
import {
	readingHistoryAtom,
	readingHistoryLoadingAtom,
	seriesAtom,
	seriesLoadingAtom,
	useSelectedVolume,
	volumesAtom,
	volumesLoadingAtom,
} from './state/data';
import { VolumeEditor } from './VolumeEditor';
import logoUrl from './logo.svg';
import { ReadingHistory } from './ReadingHistory';

function App() {
	const app = <AppLogic />;
	return (
		<BrowserRouter>
			<div className="h-full w-full flex flex-col">
				<header className="mx-auto my-3 flex items-center">
					<img src={logoUrl} className="h-16 mr-3" />
					<h1 className="text-3xl text-center font-light">Bookmark</h1>
				</header>
				<main className="flex flex-row flex-1">
					<Routes>
						<Route path="/" element={app} />
						<Route path="/series/:seriesId/volumes" element={app} />
						<Route path="/series/:seriesId/volumes/:volumeId" element={app} />
					</Routes>
				</main>
			</div>
		</BrowserRouter>
	);
}

function AppLogic() {
	const { seriesId, volumeId } = useParams(),
		setSeries = useAtom(seriesAtom)[1],
		setVolumes = useAtom(volumesAtom)[1],
		volume = useSelectedVolume(),
		setReadingHistory = useAtom(readingHistoryAtom)[1],
		setSeriesLoading = useAtom(seriesLoadingAtom)[1],
		setVolumesLoading = useAtom(volumesLoadingAtom)[1],
		setReadingHistoryLoading = useAtom(readingHistoryLoadingAtom)[1];

	useEffect(() => {
		setVolumes([]);
		setVolumesLoading(true);
		setReadingHistory([]);

		if (seriesId) {
			apiRequest(`/series/${seriesId}/volumes`).then((volumes) => {
				setVolumes(volumes);
				setVolumesLoading(false);
			});
		}
	}, [seriesId]);

	useEffect(() => {
		setSeriesLoading(true);
		apiRequest(`/series`).then((series) => {
			setSeries(series);
			setSeriesLoading(false);
		});
	}, []);

	useEffect(() => {
		setReadingHistoryLoading(true);
		setReadingHistory([]);

		apiRequest(`/series/${seriesId}/volumes/${volumeId}/history`).then((history) => {
			setReadingHistoryLoading(false);
			setReadingHistory(history);
		});
	}, [volume]);

	return (
		<div className="flex flex-1">
			<SeriesSelector />
			{seriesId && <VolumeSelector />}
			{volumeId && (
				<Fragment key={volumeId}>
					<VolumeEditor />
					<ReadingHistory />
				</Fragment>
			)}
		</div>
	);
}

export default App;
