import React from 'react';
import ReactDOM from 'react-dom';
import AuthApp from './AuthApp';
import '../app/index.css';
import './firebase-auth';

ReactDOM.render(
	<React.StrictMode>
		<AuthApp />
	</React.StrictMode>,
	document.getElementById('root')
);
