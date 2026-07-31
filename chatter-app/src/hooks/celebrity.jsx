export default async function celebrity(from, to) {
  if (from == "") {
    from = "2001-01-01"
  }
  if (to == "") {
    to = new Date().toISOString().split("T")[0];
  }

  const res = await fetch(
    `http://localhost:8000/api/celebrity?from=${from}&to=${to}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get celebrity api");
  }

  return await res.json();
}