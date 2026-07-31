export default async function bestFriend(from, to) {
  if (from == "") {
    from = "2001-01-01"
  }
  if (to == "") {
    to = new Date().toISOString().split("T")[0];
  }

  const res = await fetch(
    `http://localhost:8000/api/biggest-fan?from=${from}&to=${to}`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get biggest fan api");
  }

  return await res.json();
}