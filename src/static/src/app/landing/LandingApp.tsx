export default function LandingApp() {
	return (
		<div className="flex flex-col mx-auto">
			<div className="text-center flex flex-col">
				<p className="text-5xl my-10 font-light">Read without losing the plot</p>
				<p className="text-lg mx-auto mb-10 w-96">
					Well Read helps you organize your notes about books you're reading, so you're never lost when starting a new
					volume.
				</p>
				<p className="my-10">
					<a
						href="/auth/"
						className="p-3 bg-sky-400 rounded text-slate-800 text-2xl hover:bg-sky-300 hover:text-black transition-colors"
					>
						Login or Sign up
					</a>
				</p>
			</div>

			<div className="my-10">
				<p className="text-center text-slate-400">
					Well Read is a work in progress by{' '}
					<a className="text-sky-400 hover:underline" href="https://github.com/sheodox">
						sheodox
					</a>
					.
				</p>
			</div>
		</div>
	);
}
