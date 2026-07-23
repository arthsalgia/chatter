export default async function mostTextedDate(from, to, n) {
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
    `http://localhost:8000/most-texted-date?from=${from}&to=${to}&n=${n}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get most texted date api");
  }

  return await res.json();
}