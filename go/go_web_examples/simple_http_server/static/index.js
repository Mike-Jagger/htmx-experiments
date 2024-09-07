const scanner = new Instascan.Scanner({
	video: document.getElementById("preview"),
	scanPeriod: 5,
	refreshTime: 5,
});

let cameras_;

scanner.addListener("scan", function (content) {
	alert("Scanned QR code: " + content);
});

Instascan.Camera.getCameras()
	.then(function (cameras) {
		if (cameras.length > 0) {
			cameras_ = cameras;
		} else {
			console.error("No cameras found.");
			alert("No Cameras found.");
		}
	})
	.catch(function (e) {
		console.error(e);
	});

document.getElementById("scan-btn").addEventListener("click", function () {
	if (
		/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
			navigator.userAgent
		)
	) {
		window.location = "https://www.scan.me";
	} else {
		scanner.start(cameras_[1]);
	}
});
