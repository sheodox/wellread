import { atom, useAtom } from 'jotai';
import { useParams } from 'react-router-dom';

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

export interface ReadingHistory {
	id: number;
	volumeId: number;
	currentPage: number;
	createdAt: string;
}

export const seriesAtom = atom<Series[]>([]);
export const seriesLoadingAtom = atom<boolean>(false);
export const volumesAtom = atom<Volume[]>([]);
export const volumesLoadingAtom = atom<boolean>(false);
export const readingHistoryAtom = atom<ReadingHistory[]>([]);
export const readingHistoryLoadingAtom = atom<boolean>(false);

export function useSelectedSeries() {
	const { seriesId } = useParams(),
		[series] = useAtom(seriesAtom);

	if (!seriesId) {
		return null;
	}

	return series.find((s) => s.id === +seriesId) ?? null;
}

export function useSelectedVolume() {
	const { volumeId } = useParams(),
		[volumes] = useAtom(volumesAtom);

	if (!volumeId) {
		return null;
	}

	return volumes.find((v) => v.id === +volumeId) ?? null;
}
