import express from "express";

const app = express();

app.use(express.static("public"));
app.use(express.urlencoded({ extended: true }));
app.use(express.json());

// GET Users
app.get("/users", async (req, res) => {
	// const users = [
	// 	{ id: 1, name: "John Doe" },
	// 	{ id: 2, name: "Sherlok Holmes" },
	// 	{ id: 3, name: "Mary Jane" },
	// ];

	const response = await fetch("https://jsonplaceholder.typicode.com/users");
	const users = await response.json();

	setTimeout(async () => {
		res.status(200).send(`
			<h1 class="text-2xl font-bold my-4"> Users </h1>
			<ul>
				${users
					.map(
						(user) =>
							`<li class='mx-0 my-0'>${user.id}: ${user.name}</li>`
					)
					.join("")}
			</ul>
		`);
	}, 2000);
});

app.listen(3000, () => {
	console.log("Sever listening to port 3000");
});
