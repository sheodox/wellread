import create from 'zustand';
import { apiRequest } from './api';

export interface Series {
	id: number;
	name: string;
	notes: string;
	createdAt: string;
}

export interface Volume {
	id: number;
	seriesId: number;
	name: string;
	notes: string;
	createdAt: string;
	currentPage: number;
}

export type VolumeUpdateable = Pick<Volume, 'name' | 'notes' | 'currentPage'>;

export interface ReadingHistory {
	id: number;
	volumeId: number;
	currentPage: number;
	createdAt: string;
}

interface WellReadState {
	// null as a third state for when we don't yet know, we don't want to show
	// the app until we know which of the two apps to show
	loggedIn: boolean | null;
	series: Series[];
	volumes: Volume[];
	readingHistory: ReadingHistory[];
	seriesLoading: boolean;
	volumesLoading: boolean;
	readingHistoryLoading: boolean;
	selectedSeriesId: number | null;
	selectedVolumeId: number | null;
	setSelectedSeriesId: (seriesId: number | null) => void;
	setSelectedVolumeId: (volumeId: number | null) => void;
	loadSeries: () => Promise<void>;
	newSeries: (name: string) => Promise<void>;
	renameSeries: (id: number, name: string) => Promise<void>;
	deleteSeries: (id: number) => Promise<void>;
	loadVolumes: (seriesId: number) => Promise<void>;
	updateVolume: (id: number, volume: VolumeUpdateable) => Promise<void>;
	newVolume: (name: string) => Promise<void>;
	deleteVolume: (id: number) => Promise<void>;
	loadReadingHistory: (seriesId: number, volumeId: number) => Promise<void>;
	deleteReadingHistory: (historyId: number) => Promise<void>;
}

const MINUTE_MS = 1000 * 60,
	//series probably won't change as often as volumes, don't update those as often
	SERIES_POLL_TIME = MINUTE_MS * 60,
	//polling for volumes lets us update notes/pages
	VOLUME_POLL_TIME = MINUTE_MS * 10;
type Timer = ReturnType<typeof setTimeout>;
let seriesPollTimeout: Timer, volumePollTimeout: Timer;

export const useStore = create<WellReadState>((set, get) => {
	return {
		loggedIn: null,
		series: [],
		seriesLoading: false,
		volumes: [],
		volumesLoading: false,
		readingHistory: [],
		readingHistoryLoading: false,
		selectedSeries: null,
		selectedVolume: null,
		selectedSeriesId: null,
		selectedVolumeId: null,
		setSelectedSeriesId: (seriesId: number | null) => {
			set({ selectedSeriesId: seriesId, volumes: [] });
		},
		setSelectedVolumeId: (volumeId: number | null) => {
			set({ selectedVolumeId: volumeId, readingHistory: [] });
		},
		loadSeries: async () => {
			const load = async () => {
				set({ seriesLoading: true });
				const { body: series, status } = await apiRequest<Series[]>('/series');

				if (status === 401) {
					set({ loggedIn: false, seriesLoading: false });
				} else {
					set({ loggedIn: true, seriesLoading: false, series });
				}

				seriesPollTimeout = setTimeout(load, SERIES_POLL_TIME);
			};

			clearTimeout(seriesPollTimeout);
			await load();
		},
		newSeries: async (name: string) => {
			const { body: series } = await apiRequest<Series[]>(`/series`, 'POST', {
				name,
			});

			set({ series });
		},
		renameSeries: async (id: number, name: string) => {
			const { body: series } = await apiRequest<Series[]>(`/series/${id}`, 'PATCH', {
				name,
			});

			set({ series });
		},
		deleteSeries: async (id: number) => {
			const { body: series } = await apiRequest<Series[]>(`/series/${id}`, 'DELETE');

			set({ series });
		},
		loadVolumes: async (seriesId: number) => {
			const load = async () => {
				set({ volumesLoading: true });

				const { body: volumes } = await apiRequest<Volume[]>(`/series/${seriesId}/volumes`);
				set((state) => {
					return {
						volumesLoading: false,
						volumes,
					};
				});

				volumePollTimeout = setTimeout(load, VOLUME_POLL_TIME);
			};

			clearTimeout(volumePollTimeout);
			await load();
		},
		updateVolume: async (id: number, volume: VolumeUpdateable) => {
			const { selectedSeriesId } = get();

			const { body: volumes } = await apiRequest<Volume[]>(
				`/series/${selectedSeriesId}/volumes/${id}`,
				'PATCH',
				volume
			);
			set({ volumes: volumes });
		},
		newVolume: async (name: string) => {
			const { selectedSeriesId } = get();

			const { body: volumes } = await apiRequest<Volume[]>(`/series/${selectedSeriesId}/volumes`, 'POST', { name });
			set({ volumes: volumes });
		},
		deleteVolume: async (id: number) => {
			const { selectedSeriesId } = get();

			const { body: volumes } = await apiRequest<Volume[]>(`/series/${selectedSeriesId}/volumes/${id}`, 'DELETE');
			set({ volumes: volumes });
		},
		loadReadingHistory: async (seriesId: number, volumeId: number) => {
			set({ readingHistoryLoading: true, readingHistory: [] });

			const { body: readingHistory } = await apiRequest<ReadingHistory[]>(
				`/series/${seriesId}/volumes/${volumeId}/history`
			);
			set({ readingHistoryLoading: false, readingHistory });
		},
		deleteReadingHistory: async (historyId: number) => {
			const { selectedVolumeId, selectedSeriesId } = get();
			set({ readingHistoryLoading: true });

			const { body: history } = await apiRequest<ReadingHistory[]>(
				`/series/${selectedSeriesId}/volumes/${selectedVolumeId}/history/${historyId}`,
				'DELETE'
			);
			set({ readingHistoryLoading: false, readingHistory: history });
		},
	};
});

export function useSelectedVolume() {
	const { volumes, selectedVolumeId } = useStore();
	return volumes.find((v) => v.id == selectedVolumeId) ?? null;
}
export function useSelectedSeries() {
	const { series, selectedSeriesId } = useStore();
	return series.find((s) => s.id == selectedSeriesId) ?? null;
}
