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

	const limit = +req.query.limit || 10;

	const response = await fetch(
		`https://jsonplaceholder.typicode.com/users?_limit=${limit}`
	);
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

app.get("/calculate", async (req, res) => {
	const temp = parseFloat(+req.query.fahrenheit);

	res.status(200).send(`
		<span
			id="celcius"
			class="text-center bg-neutral-800 text-2xl font-bold w-full text-white"
		>
			${(temp - 32) * (5 / 9).toFixed(0)}
		</span>`);
});

app.get("/get-temperature", async (req, res) => {
	const randomTemp = Math.random() * 100;

	setTimeout(() => {
		res.status(200).send(`
				<p
					class="text-white text-3xl font-bold mt-5"
					hx-get="/get-temperature"
					hx-target="innerHTML"
					hx-trigger="every 5s"
					hx-indicator="#loading"
				>
					${randomTemp.toFixed(2)} degrees
				</p>
			`);
	}, 2000);
});

app.post("/search", (req, res) => {
	const search = req.body.search?.toLowerCase();
	const contacts = [
		{ email: "johndoe@email.com", name: "John Doe" },
		{ email: "sherlokholmes@email.com", name: "Sherlok Holmes" },
		{ email: "maryjane@email.com", name: "Mary Jane" },
	];

	if (!search) {
		res.send("<tr></tr>");
	}

	const searchResults = contacts.filter((contact) => {
		const name = contact.name.toLowerCase();
		const email = contact.email.toLowerCase();

		return name.includes(search) || email.includes(search);
	});

	res.status(200).send(
		searchResults
			.map(
				(contact) =>
					`<tr>
						<td><div class="my-4 p-2">${contact.name}</div></td>
						<td><div class="my-4 p-2">${contact.name}</div></td>
					<tr/>`
			)
			.join("")
	);
});

app.listen(3000, () => {
	console.log("Sever listening to port 3000");
});
