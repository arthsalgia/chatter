export default async function topMessages(from, to, n) {
  console.log(from)
  console.log(to)
  console.log(n)
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
    `http://localhost:8000/nth-common?from=${from}&to=${to}&n=${n}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get nth common api");
  }

  return await res.json();
}