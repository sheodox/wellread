import { Link, useNavigate } from 'react-router-dom';
import { Spinner } from './Spinner';
import { createPopper } from '@popperjs/core';
import { useSelectedSeries, useSelectedVolume, Volume, useStore } from './state/data';
import { DotsVerticalIcon } from '@heroicons/react/outline';
import { ComponentType, useLayoutEffect, useRef, useState } from 'react';
import { Empty } from './Empty';
import { theme } from './theme';
import Portal from './Portal';
import StatusBadge from './StatusBadge';

interface SelectorItem {
	id: number;
	name: string;
}

type HrefGenerator = (item: SelectorItem) => string;

interface SelectorPropsOptions {
	promptQuestion: string;
	onNew: (name: string) => any;
	onDelete: (item: SelectorItem) => any;
	onRename: (item: SelectorItem, name: string) => any;
}
type SelectorProps = {
	title: string;
	loading: boolean;
	href: HrefGenerator;
	selectedId?: number | null;
	items: SelectorItem[];
	itemRenderer?: ComponentType<{ item: SelectorItem }>;
};

export function EditableSelector(props: SelectorProps & SelectorPropsOptions) {
	const promptNew = () => {
		const name = prompt(props.promptQuestion)?.trim();
		if (name) {
			props.onNew(name);
		}
	};

	return (
		<div className="w-sm">
			<div className="flex justify-between border-b border-slate-700 p-4 items-center mb-6 gap-4">
				<h1 className="text-3xl flex items-baseline">
					{props.title}
					<div className="ml-3 flex justify-center">
						<Spinner show={props.loading} />
					</div>
				</h1>
				<button className={theme.button.secondary} onClick={promptNew}>
					Add
				</button>
			</div>
			{!props.loading && props.items.length === 0 && <Empty />}
			<ul>
				{props.items.map((item) => {
					return (
						<SelectorListItem
							key={item.id}
							item={item}
							selectedId={props.selectedId}
							href={props.href}
							onRename={props.onRename}
							onDelete={props.onDelete}
							itemRenderer={props.itemRenderer}
						/>
					);
				})}
			</ul>
		</div>
	);
}
export function ReadonlySelector(props: SelectorProps) {
	return (
		<div className="w-sm">
			<div className="flex justify-between border-b border-slate-700 p-4 items-center mb-6 gap-4">
				<h1 className="text-3xl flex items-baseline">
					{props.title}
					<div className="ml-3 flex justify-center">
						<Spinner show={props.loading} />
					</div>
				</h1>
			</div>
			{!props.loading && props.items.length === 0 && <Empty />}
			<ul>
				{props.items.map((item) => {
					return (
						<li className={`flex rounded-md border-2 border-transparent hover:border-sky-600`} key={item.id}>
							<Link className={`flex-1 p-4 font-bold hover:text-white text-slate-400`} to={props.href(item)}>
								{props.itemRenderer ? <props.itemRenderer item={item} /> : item.name}
							</Link>
						</li>
					);
				})}
			</ul>
		</div>
	);
}

function SelectorListItem(props: {
	href: HrefGenerator;
	selectedId?: number | null;
	item: SelectorItem;
	onDelete: SelectorPropsOptions['onDelete'];
	onRename: SelectorPropsOptions['onRename'];
	itemRenderer?: ComponentType<{ item: SelectorItem }>;
}) {
	const active = props.selectedId == props.item.id,
		menu = useRef(null),
		menuTrigger = useRef(null),
		[showMenu, setShowMenu] = useState(false);

	function openMenu(e: React.MouseEvent) {
		// need to stop propagation now, the menu opens and the destroy handler is
		// bound synchronously, without this it'll close immediately
		e.stopPropagation();
		setShowMenu(!showMenu);
	}

	useLayoutEffect(() => {
		if (showMenu && menuTrigger.current && menu.current) {
			const popper = createPopper(menuTrigger.current, menu.current, {
				placement: 'bottom',
			});
			document.body.addEventListener(
				'click',
				() => {
					popper.destroy();
					setShowMenu(false);
				},
				{ once: true }
			);
		}
	}, [showMenu]);

	return (
		<li className={`flex rounded-md border-2 border-transparent hover:border-sky-600 ${active ? 'bg-slate-700' : ''}`}>
			<Link
				className={`flex-1 p-4 font-bold hover:text-white ${active ? 'text-white' : 'text-slate-400'}`}
				to={props.href(props.item)}
			>
				{props.itemRenderer ? <props.itemRenderer item={props.item} /> : props.item.name}
			</Link>
			<button onClick={openMenu} ref={menuTrigger}>
				<DotsVerticalIcon className="h-5 px-2 hover:text-sky-400 transition-colors" />
				<span className="sr-only">Open Menu</span>
			</button>
			{showMenu && (
				<Portal>
					<div ref={menu}>
						<SelectorMenu onDelete={props.onDelete!} onRename={props.onRename!} item={props.item} />
					</div>
				</Portal>
			)}
		</li>
	);
}

function SelectorMenu(props: { item: SelectorItem } & Pick<SelectorPropsOptions, 'onRename' | 'onDelete'>) {
	const promptRename = () => {
			const name = prompt('Enter a new name', props.item.name)?.trim();
			if (name) {
				props.onRename(props.item, name);
			}
		},
		confirmDelete = () => {
			if (confirm(`Are you sure you want to delete ${props.item.name}?`)) {
				props.onDelete(props.item);
			}
		},
		items = [
			{
				text: 'Rename',
				handler: promptRename,
			},
			{ text: 'Delete', handler: confirmDelete },
		];

	return (
		<div className="bg-slate-700 rounded overflow-hidden">
			<ul>
				{items.map((item, i) => {
					return (
						<li key={i} className="hover:bg-gray-600">
							<button className="p-3 font-bold" onClick={item.handler}>
								{item.text}
							</button>
						</li>
					);
				})}
			</ul>
		</div>
	);
}

export function SeriesSelector() {
	const { seriesLoading, series, newSeries, renameSeries, deleteSeries } = useStore(),
		navigate = useNavigate(),
		selectedSeries = useSelectedSeries(),
		seriesHref = (series: SelectorItem) => `/series/${series.id}/volumes`,
		onRename = async (series: SelectorItem, name: string) => {
			await renameSeries(series.id, name);
		},
		onDelete = async (series: SelectorItem) => {
			if (series.id == selectedSeries?.id) {
				navigate('/');
			}
			await deleteSeries(series.id);
		};

	return (
		<EditableSelector
			loading={seriesLoading}
			onNew={newSeries}
			promptQuestion="Enter a new series name"
			title="Series"
			href={seriesHref}
			items={series}
			selectedId={selectedSeries?.id ?? null}
			onRename={onRename}
			onDelete={onDelete}
		></EditableSelector>
	);
}

export function VolumeItem(props: { item: SelectorItem }) {
	const volume = props.item as Volume;
	return (
		<div className="flex items-center gap-2">
			<StatusBadge status={volume.status} />
			{volume.name}
		</div>
	);
}

export function ReadingVolumeItem(props: { item: SelectorItem }) {
	const volume = props.item as Volume,
		{ series } = useStore(),
		seriesName = series.find((s) => s.id === volume.seriesId)?.name;

	return (
		<div className="flex items-center gap-2">
			<StatusBadge status={volume.status} />
			{volume.name}
			{seriesName ? ` - ${seriesName}` : ''}
		</div>
	);
}

export function VolumeSelector() {
	const { selectedVolumeId, volumes, volumesLoading, newVolume, deleteVolume, updateVolume } = useStore(),
		navigate = useNavigate(),
		selectedSeries = useSelectedSeries(),
		selectedVolume = useSelectedVolume(),
		seriesHref = (volume: SelectorItem) => `/series/${selectedSeries?.id}/volumes/${volume.id}`,
		onRename = async (volume: SelectorItem, name: string) => {
			const v = volume as Volume;
			await updateVolume(volume.id, {
				name,
				currentPage: v.currentPage,
				notes: v.notes,
				status: v.status,
			});
		},
		onDelete = async (volume: SelectorItem) => {
			if (volume.id === selectedVolumeId) {
				navigate(`/series/${selectedSeries?.id}/volumes`);
			}
			await deleteVolume(volume.id);
		};

	return (
		<EditableSelector
			loading={volumesLoading}
			onNew={newVolume}
			promptQuestion="Enter a new volume name"
			title="Volumes"
			href={seriesHref}
			items={volumes}
			selectedId={selectedVolume?.id ?? null}
			onRename={onRename}
			onDelete={onDelete}
			itemRenderer={VolumeItem}
		></EditableSelector>
	);
}

export function ReadingVolumesSelector() {
	const { readingVolumes, readingVolumesLoading } = useStore(),
		seriesHref = (item: SelectorItem) => {
			const volume = item as Volume;
			return `/series/${volume.seriesId}/volumes/${volume.id}`;
		};

	return (
		<ReadonlySelector
			loading={readingVolumesLoading}
			title="Reading"
			href={seriesHref}
			items={readingVolumes}
			itemRenderer={ReadingVolumeItem}
		></ReadonlySelector>
	);
}
