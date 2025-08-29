import { fetchAuthSession } from 'aws-amplify/auth';

export async function getIdToken(): Promise<string | null> {
  try {
    const session = await fetchAuthSession();

    const idToken = session.tokens?.idToken?.toString() ?? null;
    return idToken;
  } catch (e) {
    return null;
  }
}
