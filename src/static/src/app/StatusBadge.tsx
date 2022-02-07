const statusColorMapping: Record<string, string> = {
	planning: 'bg-orange-400',
	reading: 'bg-blue-400',
	completed: 'bg-green-400',
	dropped: 'bg-red-400',
};

function getStatusColor(status: string) {
	return statusColorMapping[status];
}

export default function StatusBadge(props: { status: string; size?: 'large' | 'small' }) {
	const size = props.size === 'large' ? 'h-4 w-4' : 'h-2 w-2';
	return <div className={`rounded-full ${size} ${getStatusColor(props.status)}`} title={props.status} />;
}
