export const collectFormValues = (form: FormData) => {
	// eslint-disable-next-line @typescript-eslint/no-explicit-any
	const data: Record<string, any> = {};

	for (const [key, value] of form.entries()) {
		data[key] = value;
	}

	return data;
};
