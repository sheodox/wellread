import logoUrl from './logo.svg';

export default function Header() {
	return (
		<header className="inline-block my-3 flex items-center">
			<img src={logoUrl} className="h-16 mr-3" />
			<h1 className="text-3xl text-center font-light">Well Read</h1>
		</header>
	);
}
