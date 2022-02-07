import { Fragment, useEffect, useState } from 'react';
import { BrowserRouter, Link, Route, Routes, useParams } from 'react-router-dom';
import { ReadingVolumesSelector, SeriesSelector, VolumeSelector } from './Selector';
import { useSelectedSeries, useSelectedVolume, useStore } from './state/data';
import { VolumeEditor } from './VolumeEditor';
import { ReadingHistory } from './ReadingHistory';
import Header from './Header';
import Footer from './Footer';
import LandingApp from './landing/LandingApp';

function App() {
	const app = <AppLogic />;
	return (
		<BrowserRouter>
			<div className="h-full w-full flex flex-col">
				<div className="mx-auto">
					<Link to="/">
						<Header />
					</Link>
				</div>
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
		{
			loggedIn,
			loadReading,
			readingVolumes,
			loadVolumes,
			loadSeries,
			loadReadingHistory,
			setSelectedSeriesId,
			setSelectedVolumeId,
		} = useStore(),
		selectedVolume = useSelectedVolume(),
		selectedSeries = useSelectedSeries(),
		// this is used to change the page justification so after the user has selected a series we don't try and center the page,
		// causing jumps when switching series.
		[hasSelectedSomething, setHasSelectedSomething] = useState(false);

	useEffect(() => {
		setSelectedSeriesId(seriesId ? +seriesId : null);

		if (seriesId) {
			loadVolumes(parseInt(seriesId, 10));
		}
	}, [seriesId]);

	useEffect(() => {
		loadReading();
		loadSeries();
	}, []);

	useEffect(() => {
		setSelectedVolumeId(volumeId ? +volumeId : null);
		if (volumeId) {
			setHasSelectedSomething(true);
		}
	}, [volumeId]);

	useEffect(() => {
		if (selectedSeries && selectedVolume) {
			loadReadingHistory(selectedSeries.id, selectedVolume.id);
		}
	}, [selectedSeries, selectedVolume]);

	return (
		<>
			{loggedIn === true && (
				<div className="flex flex-1 flex-col max-w-screen-2xl mx-auto">
					<div
						className={`px-4 flex flex-1 flex-col gap-14 md:flex-row ${hasSelectedSomething ? '' : 'justify-center'}`}
					>
						<div className="flex flex-col gap-14">
							{readingVolumes.length !== 0 && <ReadingVolumesSelector />}
							<SeriesSelector />
						</div>
						{seriesId && <VolumeSelector />}
						{volumeId && (
							<Fragment key={selectedVolume?.id}>
								<VolumeEditor />
								<ReadingHistory />
							</Fragment>
						)}
					</div>
					<Footer />
				</div>
			)}
			{loggedIn === false && <LandingApp />}
		</>
	);
}

export default App;
