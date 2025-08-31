export type SnakeCase<T> = T extends Record<string, unknown>
  ? { [K in keyof T as K extends string ? SnakeCaseKey<K> : K]: SnakeCase<T[K]> }
  : T extends (infer U)[]
    ? U extends Record<string, unknown>
      ? SnakeCase<U>[]
      : T
    : T;

type SnakeCaseKey<S> = S extends `${infer T}${infer U}`
  ? U extends Capitalize<U>
    ? `${Lowercase<T>}_${SnakeCaseKey<Uncapitalize<U>>}`
    : `${Lowercase<T>}${SnakeCaseKey<U>}`
  : S;

function toSnakeCase<T>(obj: T): SnakeCase<T> {
  if (Array.isArray(obj)) {
    return obj.map((v) => toSnakeCase(v)) as unknown as SnakeCase<T>;
  }
  if (obj !== null && typeof obj === 'object') {
    return Object.fromEntries(
      Object.entries(obj as Record<string, unknown>).map(([key, value]) => [
        key.replace(/[A-Z]/g, (letter) => `_${letter.toLowerCase()}`),
        toSnakeCase(value),
      ]),
    ) as SnakeCase<T>;
  }
  return obj as SnakeCase<T>;
}

export default toSnakeCase;
