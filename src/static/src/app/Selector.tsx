import { Link, useNavigate } from 'react-router-dom';
import { Spinner } from './Spinner';
import { createPopper } from '@popperjs/core';
import { useSelectedSeries, useSelectedVolume, Volume, useStore } from './state/data';
import { DotsVerticalIcon } from '@heroicons/react/outline';
import { useEffect, useLayoutEffect, useRef, useState } from 'react';
import { Empty } from './Empty';
import { theme } from './theme';

interface SelectorItem {
	id: number;
	name: string;
}

type HrefGenerator = (item: SelectorItem) => string;

interface SelectorProps {
	title: string;
	loading: boolean;
	href: HrefGenerator;
	selectedId: number | null;
	items: SelectorItem[];
	promptQuestion: string;
	onNew: (name: string) => any;
	onDelete: (item: SelectorItem) => any;
	onRename: (item: SelectorItem, name: string) => any;
}

export function Selector(props: SelectorProps) {
	const promptNew = () => {
		const name = prompt(props.promptQuestion)?.trim();
		if (name) {
			props.onNew(name);
		}
	};

	return (
		<div className="w-sm mx-9">
			<div className="flex justify-between border-b border-slate-700 p-4 items-center mb-6">
				<h1 className="text-3xl">{props.title}</h1>
				<button className={`ml-6 ${theme.button.secondary}`} onClick={promptNew}>
					Add
				</button>
			</div>
			{props.loading && (
				<div className="flex justify-center">
					<Spinner />
				</div>
			)}
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
						/>
					);
				})}
			</ul>
		</div>
	);
}

function SelectorListItem(props: {
	href: HrefGenerator;
	selectedId: number | null;
	item: SelectorItem;
	onRename: SelectorProps['onRename'];
	onDelete: SelectorProps['onDelete'];
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
		<li className="flex">
			<Link
				className={`flex-1 p-4 rounded-md font-bold hover:text-sky-400 transition-colors ${
					active ? 'text-sky-400 drop-shadow' : 'text-slate-400'
				}`}
				to={props.href(props.item)}
			>
				{props.item.name}
			</Link>
			<button onClick={openMenu} ref={menuTrigger}>
				<DotsVerticalIcon className="h-5 px-2 hover:text-sky-400 transition-colors" />
			</button>
			{showMenu && (
				<div ref={menu}>
					<SelectorMenu onDelete={props.onDelete} onRename={props.onRename} item={props.item} />
				</div>
			)}
		</li>
	);
}

function SelectorMenu(props: { item: SelectorItem } & Pick<SelectorProps, 'onRename' | 'onDelete'>) {
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
		<div className="bg-slate-800 rounded overflow-hidden">
			<ul>
				{items.map((item, i) => {
					return (
						<li key={i} className="hover:text-sky-200 hover:bg-gray-700">
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
		<Selector
			loading={seriesLoading}
			onNew={newSeries}
			promptQuestion="Enter a new series name"
			title="Series"
			href={seriesHref}
			items={series}
			selectedId={selectedSeries?.id ?? null}
			onRename={onRename}
			onDelete={onDelete}
		></Selector>
	);
}

export function VolumeSelector() {
	const { selectedVolumeId, volumes, volumesLoading, newVolume, deleteVolume, updateVolume } = useStore(),
		navigate = useNavigate(),
		selectedSeries = useSelectedSeries(),
		selectedVolume = useSelectedVolume(),
		seriesHref = (volume: SelectorItem) => `/series/${selectedSeries?.id}/volumes/${volume.id}`,
		onRename = async (volume: SelectorItem, name: string) => {
			await updateVolume(volume.id, {
				name,
				currentPage: (volume as Volume).currentPage,
				notes: (volume as Volume).notes,
			});
		},
		onDelete = async (volume: SelectorItem) => {
			if (volume.id === selectedVolumeId) {
				navigate(`/series/${selectedSeries?.id}/volumes`);
			}
			await deleteVolume(volume.id);
		};

	return (
		<Selector
			loading={volumesLoading}
			onNew={newVolume}
			promptQuestion="Enter a new volume name"
			title="Volumes"
			href={seriesHref}
			items={volumes}
			selectedId={selectedVolume?.id ?? null}
			onRename={onRename}
			onDelete={onDelete}
		></Selector>
	);
}
