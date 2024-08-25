import express from "express";
import cors from "cors";
import xss from "xss";

const app = express();

const corsOptions = {
	credentials: true,
	origin: ["http:localhost:3000"],
};

app.use(cors(corsOptions));
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
		{ name: "John Doe", email: "john@example.com" },
		{ name: "Jane Doe", email: "jane@example.com" },
		{ name: "Alice Smith", email: "alice@example.com" },
		{ name: "Bob Williams", email: "bob@example.com" },
		{ name: "Mary Harris", email: "mary@example.com" },
		{ name: "David Mitchell", email: "david@example.com" },
	];

	if (!search) {
		res.status(200).send("<tr></tr>");
		return;
	}

	const searchResults = contacts.filter((contact) => {
		const name = contact.name.toLowerCase();
		const email = contact.email.toLowerCase();

		return name.includes(search) || email.includes(search);
	});

	setTimeout(() => {
		res.status(200).send(
			searchResults
				.map(
					(contact) =>
						`<tr>
						<td><div class="my-4 p-2">${contact.name}</div></td>
						<td><div class="my-4 p-2">${contact.email}</div></td>
					<tr/>`
				)
				.join("")
		);
	}, 2000);
});

app.post("/search/api", async (req, res) => {
	const search = req.body.search?.toLowerCase();

	const response = await fetch(`https://jsonplaceholder.typicode.com/users`);
	const contacts = await response.json();

	if (!search) {
		res.status(200).send("<tr></tr>");
		return;
	}

	const searchResults = contacts.filter((contact) => {
		const name = contact.name.toLowerCase();
		const email = contact.email.toLowerCase();

		return name.includes(search) || email.includes(search);
	});

	setTimeout(() => {
		res.status(200).send(
			searchResults
				.map(
					(contact) =>
						`<tr>
						<td><div class="my-4 p-2">${contact.name}</div></td>
						<td><div class="my-4 p-2">${contact.email}</div></td>
					<tr/>`
				)
				.join("")
		);
	}, 1000);
});

app.post("/contact/email", (req, res) => {
	const email = req.body.email;

	const emailRegex = /^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}$/;

	const isValid = {
		message: "Email is valid",
		class: "text-green-700",
	};

	const isInValid = {
		message: "Email is invalid",
		class: "text-red-700",
	};

	if (emailRegex.test(email)) {
		res.send(`
			<div class="mb-4" hx-target="this" hx-swap="outerHTML">
				<label
					class="block text-gray-700 text-sm font-bold mb-2"
					for="email"
					>Email Address</label
				>
				<input
					name="email"
					hx-post="/contact/email"
					class="border rounded-lg py-2 px-3 w-full focus:outline-none focus:border-blue-500"
					type="email"
					id="email"
					value="${email}"
					required
				/>
				<div class='${isValid.class}'>${isValid.message}</div>
			</div>	
		`);
	} else {
		res.send(`
			<div class="mb-4" hx-target="this" hx-swap="outerHTML">
				<label
					class="block text-gray-700 text-sm font-bold mb-2"
					for="email"
					>Email Address</label
				>
				<input
					name="email"
					hx-post="/contact/email"
					class="border rounded-lg py-2 px-3 w-full focus:outline-none focus:border-blue-500"
					type="email"
					id="email"
					value="${email}"
					required
				/>
				<div class='${isInValid.class}'>${isInValid.message}</div>
			</div>	
		`);
	}
});

app.get("/profile/:id/edit", (req, res) => {
	const profileId = req.url.id;
	res.status(200).send(
		`
		<div
			class="container mx-auto py-8 max-w-lg"
			hx-target="this"
			hx-swap="outerHTML"
		>
			<form hx-put="/profile/${profileId}" hx-target="this" hx-swap="outerHTML">
				<div class="bg-white p-6 rounded-lg shadow-lg">
				<div class="mb-4">
					<label for="name" class="text-lg font-semibold">Name</label>
					<input type="text" id="name" name="name" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring focus:ring-blue-400" value="John Doe">
				</div>
				<div class="mb-4">
					<label for="bio" class="text-lg font-semibold">Bio</label>
					<textarea id="bio" name="bio" class="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring focus:ring-blue-400" rows="6">Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla vestibulum vestibulum elit, ac facilisis ipsum eleifend sed. Duis tincidunt augue nec neque cursus, nec aliquet purus tempor.</textarea>
				</div>
				<div class="mt-6">
				<button type="submit" class="px-4 py-2 bg-indigo-700 text-white rounded-lg hover:bg-indigo-600">Save Changes</button>
				<button type="button" hx-get="/index.html" class="px-4 py-2 bg-gray-300 text-gray-700 rounded-lg ml-2 hover:bg-gray-400">Cancel</button>
				</div>
				</div>
			</form>
		</div>
		`
	);
});

app.put("/profile/:id", (req, res) => {
	const name = req.body.name;
	const bio = req.body.bio;

	res.status(200).send(
		`
		<div
			class="container mx-auto py-8 max-w-lg"
			hx-target="this"
			hx-swap="outerHTML"
		>
			<div class="bg-white p-6 rounded-lg shadow-lg">
				<h2 class="text-xl text-gray-400 mb-4">PROFILE</h2>
				<h1 class="text-2xl font-bold mb-4">${name}</h1>
				<p class="text-gray-700">${bio}</p>
				<button
					hx-get="/profile/1/edit"
					class="bg-indigo-700 text-white font-bold w-full py-2 px-4 rounded-lg mt-4 hover:bg-indigo-600"
				>
					Click To Edit
				</button>
			</div>
		</div>
		`
	);
});

app.listen(3000, () => {
	console.log("Sever listening to port 3000");
});
