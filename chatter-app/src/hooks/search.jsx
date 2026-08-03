export default async function search(from, to, word, partial) {
  if (from == "") {
    from = "2001-01-01"
  }
  if (to == "") {
    to = new Date().toISOString().split("T")[0];
  }
  if (word == "") {
    return
  }

  const res = await fetch(
    `http://localhost:8000/api/search?from=${from}&to=${to}&word=${word}&partial=${partial}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get most texted date api");
  }

  return await res.json();
}