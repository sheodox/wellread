import { useAtom } from 'jotai';
import { useEffect } from 'react';
import { BrowserRouter, Route, Routes, useParams } from 'react-router-dom';
import { SeriesSelector, VolumeSelector } from './Selector';
import { apiRequest } from './state/api';
import { seriesAtom, seriesLoadingAtom, volumesAtom, volumesLoadingAtom } from './state/series';
import { VolumeEditor } from './VolumeEditor';
import logoUrl from './logo.svg';

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
		setSeriesLoading = useAtom(seriesLoadingAtom)[1],
		setVolumesLoading = useAtom(volumesLoadingAtom)[1];

	useEffect(() => {
		setVolumes([]);
		setVolumesLoading(true);
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

	return (
		<div className="flex flex-1">
			<SeriesSelector />
			{seriesId && <VolumeSelector />}
			{volumeId && <VolumeEditor key={volumeId} />}
		</div>
	);
}

export default App;
