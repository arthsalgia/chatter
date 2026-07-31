export default async function mostCommonWord(from, to, n) {
  if (from == "") {
    from = "2001-01-01"
  }
  if (to == "") {
    to = new Date().toISOString().split("T")[0];
  }
  if (n == "") {
    n = 3
  }

  const res = await fetch(
    `http://localhost:8000/api/most-common-word?from=${from}&to=${to}&n=${n}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get nth common api");
  }

  return await res.json();
}