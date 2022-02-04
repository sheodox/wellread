import { Component, PropsWithChildren } from 'react';
import { createPortal } from 'react-dom';

export default class Portal extends Component {
	el: HTMLDivElement;
	constructor(props: PropsWithChildren<{}>) {
		super(props);
		this.el = document.createElement('div');
	}

	componentDidMount() {
		document.body.appendChild(this.el);
	}

	componentWillUnmount() {
		document.body.removeChild(this.el);
	}

	render() {
		return createPortal(this.props.children, this.el);
	}
}
