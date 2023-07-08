"use strict";

const connectedElemented = document.querySelector("#connected");
const monitorButton = document.querySelector("#monitor-button");
const dataElement = document.querySelector("#data");
const popupContainer = document.querySelector("#popup-container");
const connectedStream = new EventSource("/api/get/quantity");

let selectedIp;
let dataStream;

connectedStream.addEventListener('updateDevicesCount', (event) => {
	connectedElemented.textContent = event.data + " devices currently connected .";
});

const setUser = async (e) => {
	const userNumber = e.target.dataset.un;

	const blob = await fetch(`/api/set/user?un=${userNumber}&ip=${selectedIp}`);
	const text = await blob.text();

	if (text == "Success\n") {
		document.querySelector(".popup").firstChild.click();
	} else {
		alert("Error sending request. Check console.");
		console.log(text);
	}
};

const queryUsers = async () => {
	const searchQuery = document.querySelector("input").value;
	let baseUrl = "/api/get/users";
	if (searchQuery != "")
		baseUrl += `?s=${searchQuery}`;
	const blob = await fetch(baseUrl);
	const obj = await blob.json();

	const target = document.querySelector("#search-table");
	let children = [];

	if (!obj.Users) {
		target.replaceChildren();
		return
	}

	obj.Users.forEach( (result) => {
		const row = document.createElement("tr");

		const number = document.createElement("td");
		number.textContent = result.Number;
		row.appendChild(number);

		const name = document.createElement("td");
		name.textContent = result.Name;
		row.appendChild(name);
		
		const button_container = document.createElement("td");
		const button = document.createElement("button");
		button.textContent = "Set to this patient";
		button.dataset.un = result.Number;
		button.addEventListener('click', setUser);
		button_container.appendChild(button);
		row.appendChild(button_container);

		children.push(row);
	});

	target.replaceChildren(...children);
};

const enablePopup = (popup) => { 
	popupContainer.style.display = "flex";
	popupContainer.replaceChildren(popup);
};
const disablePopup = () => {
	popupContainer.style.display = "none";
	popupContainer.replaceChildren();
}

const createPopupElement = () => {
	const popup = document.createElement("div");
	popup.classList.add("popup");

	const close = document.createElement("span");
	close.classList.add("close-btn");
	close.classList.add("near-text");
	popup.appendChild(close);

	close.addEventListener('click', disablePopup, {"once": true});

	return popup;
}

const userAssignPopup = (e) => {
	const popup = createPopupElement();

	const h2 = document.createElement("h2");
	h2.textContent = "User assignment";
	h2.classList.add("near-text");
	popup.appendChild(h2);

	const p = document.createElement("p");
	p.textContent = `Selecting user for the IP ${e.target.dataset.ip}.`
	p.classList.add("middle-text");
	popup.appendChild(p);

	const row1 = document.createElement("div");
	row1.classList.add("row");

	const input = document.createElement("input");
	input.placeholder = "User name or number";
	input.classList.add("middle-distance");
	row1.appendChild(input);

	const searchbutton = document.createElement("button");
	searchbutton.classList.add("middle-distance");
	searchbutton.textContent = "Search";
	searchbutton.addEventListener('click', queryUsers);
	popup.firstChild.addEventListener('click', () =>
		searchbutton.removeEventListener('click', queryUsers),
		{ "once": true });
	row1.appendChild(searchbutton);

	popup.appendChild(row1);


	const table = document.createElement("table");
	const tableHeader = document.createElement("thead");
	const tableRow = document.createElement("tr");
	const numbers = document.createElement("th");
	numbers.textContent = "Numbers";
	numbers.classList.add("middle-text")
	tableRow.appendChild(numbers);
	const names = document.createElement("th");
	names.textContent = "Names";
	names.classList.add("middle-text")
	tableRow.appendChild(names);
	const actions = document.createElement("th");
	actions.textContent = "Actions";
	actions.classList.add("middle-text")
	tableRow.appendChild(actions);
	tableHeader.appendChild(tableRow);
	table.appendChild(tableHeader);

	const tableBody = document.createElement("tbody");
	tableBody.setAttribute("id", "search-table");
	table.appendChild(tableBody);

	popup.appendChild(table);

	enablePopup(popup);

	selectedIp = e.target.dataset.ip;
}

const changeOrder = async () => {
	const message = document.querySelector("input").value;

	const blob = await fetch(`/api/set/order?ip=${selectedIp}&o=${encodeURIComponent(message)}`);
	const text = await blob.text();

	if (text == "Success\n") {
		document.querySelector(".popup").firstChild.click();
	} else {
		alert("Error sending request. Check console.");
		console.log(text);
	}
}

const changeOrderPopup = (e) => {
	const popup = createPopupElement();

	const h2 = document.createElement("h2");
	h2.textContent = "Order update";
	h2.classList.add("near-text");
	popup.appendChild(h2);

	const p = document.createElement("p");
	p.textContent = `What order do you want to set to ${e.target.dataset.username}?`
	p.classList.add("middle-text");
	popup.appendChild(p);

	const row1 = document.createElement("div");
	row1.classList.add("row");

	const input = document.createElement("input");
	input.placeholder = "Go to room xxx";
	input.classList.add("middle-distance");
	row1.appendChild(input);

	const searchbutton = document.createElement("button");
	searchbutton.classList.add("middle-distance");
	searchbutton.textContent = "Update";
	searchbutton.addEventListener('click', changeOrder);
	popup.firstChild.addEventListener('click', () =>
		searchbutton.removeEventListener('click', changeOrder),
		{ "once": true });
	row1.appendChild(searchbutton);

	popup.appendChild(row1);

	enablePopup(popup);

	selectedIp = e.target.dataset.ip;
}

const processDataStream = (data) => {
	let elements = [];
	for (const ip in data) {
		const row = document.createElement('tr');

		const ipElement = document.createElement('td');
		ipElement.textContent = ip;
		ipElement.classList.add("far-text");
		row.appendChild(ipElement);

		const patientName = document.createElement('td');
		patientName.textContent = data[ip].UserNumber;
		patientName.classList.add("far-text");
		row.appendChild(patientName);

		const patientNumber = document.createElement('td');
		patientNumber.textContent = data[ip].UserName;
		patientNumber.classList.add("far-text");
		row.appendChild(patientNumber);

		const order = document.createElement('td');
		order.textContent = data[ip].Order;
		order.classList.add("far-text");
		row.appendChild(order);

		const options = document.createElement('td');
		options.classList.add("user-options");
		if (data[ip].UserNumber === 0) {
			const assign = document.createElement('button');
			assign.textContent = "Assign to a user";
			assign.classList.add("far");
			assign.addEventListener('click', userAssignPopup);
			assign.dataset.ip = ip;
			options.appendChild(assign);
		} else {
			const change = document.createElement('button');
			change.textContent = "Change user";
			change.classList.add("far");
			change.dataset.ip = ip;
			change.dataset.username = data[ip].UserName;
			change.addEventListener('click', userAssignPopup);
			options.appendChild(change);

			const order = document.createElement('button');
			order.textContent = "Change order";
			order.classList.add("far");
			order.dataset.ip = ip;
			order.addEventListener('click', changeOrderPopup);
			options.appendChild(order);
		}
		
		row.appendChild(options);

		elements.push(row);
	}
	dataElement.replaceChildren(...elements);
}

const startDataStream = () => {
	monitorButton.textContent = "Stop monitoring"
	dataStream = new EventSource("/api/get/data");
	dataStream.addEventListener('updateDevicesData', (e) => processDataStream(JSON.parse(e.data)));
	monitorButton.addEventListener('click', closeDataStream, {once: true});
}

const closeDataStream = () => {
	monitorButton.textContent = "Start monitoring"
	dataStream.close();
	monitorButton.addEventListener('click', startDataStream, {once: true});
	dataElement.replaceChildren();
}

monitorButton.addEventListener('click', startDataStream, {once: true});
