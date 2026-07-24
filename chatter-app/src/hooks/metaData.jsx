export default async function search() {


  const res = await fetch(
    `http://localhost:8000/meta-data`,
    {
      method: "GET",
    }
  );

  if (!res.ok) {
    throw new Error("Failed to get most meta data api");
  }

  return await res.json();
}