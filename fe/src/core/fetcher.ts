function constructUrl(route: string) {
  const baseUrl = import.meta.env.VITE_API_BASE_URL || '';
  return `${baseUrl}${route}`;
}

export default async function fetcher<T extends any>(
  route: string,
  options: RequestInit
): Promise<T> {
  const url = constructUrl(route);

  const response = await fetch(url, options);
  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
}
