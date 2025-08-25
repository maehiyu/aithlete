export type CamelCase<T> = T extends Record<string, unknown>
	? { [K in keyof T as K extends string ? CamelCaseKey<K> : K]: CamelCase<T[K]> }
	: T extends (infer U)[]
		? U extends Record<string, unknown>
			? CamelCase<U>[]
			: T
		: T;

type CamelCaseKey<S> = S extends `${infer T}_${infer U}`
	? `${Lowercase<T>}${Capitalize<CamelCaseKey<U>>}`
	: S extends `${infer T}-${infer U}`
		? `${Lowercase<T>}${Capitalize<CamelCaseKey<U>>}`
		: S;

function toCamelCase<T>(obj: T): CamelCase<T> {
	if (Array.isArray(obj)) {
		return obj.map((v) => toCamelCase(v)) as unknown as CamelCase<T>;
	}
	if (obj !== null && typeof obj === 'object') {
		return Object.fromEntries(
			Object.entries(obj as Record<string, unknown>).map(([key, value]) => [
				key.replace(/([-_][a-z])/g, (group) => group.toUpperCase().replace('-', '').replace('_', '')),
				toCamelCase(value),
			]),
		) as CamelCase<T>;
	}
	return obj as CamelCase<T>;
}

export default toCamelCase;
