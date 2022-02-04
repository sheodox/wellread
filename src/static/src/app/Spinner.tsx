export function Spinner(props: { show: boolean }) {
	return (
		<div
			className={`animate-spin h-4 w-4 rounded-full border-2 border-r-transparent border-sky-400 transition-opacity delay-200 ${
				props.show ? 'opacity-100' : 'opacity-0'
			}`}
		></div>
	);
}
