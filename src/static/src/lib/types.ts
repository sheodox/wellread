export interface Series {
	id: number;
	name: string;
	notes: string;
	createdAt: string;
	volumeCount: number;
}

export interface Volume {
	id: number;
	seriesId: number;
	name: string;
	notes: string;
	createdAt: string;
	status: string;
	currentPage: number;
	seriesName: string;
}

export type VolumeUpdateable = Pick<Volume, 'name' | 'notes' | 'currentPage' | 'status'> & { pagesRead?: number };

export interface ReadingHistory {
	id: number;
	volumeId: number;
	currentPage: number;
	pagesRead: number;
	createdAt: string;
}
