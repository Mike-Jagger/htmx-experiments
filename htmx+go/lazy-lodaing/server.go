package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    http.HandleFunc("/", servePage)
    http.HandleFunc("/action-buttons", serveActionButtons)
    http.HandleFunc("/image", serveImage)

    fmt.Println("Server started at :8080")
    http.ListenAndServe(":8080", nil)
}

func servePage(w http.ResponseWriter, r *http.Request) {
    page := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Lazy Load Example</title>
        <script src="https://unpkg.com/htmx.org"></script>
		<script src="https://cdn.tailwindcss.com"></script>
    </head>
    <body>
        <div id="content" class="flex flex-row items-center w-full h-[650px]">
            <!-- Action Buttons (Placeholder) -->

            <div class="bg-gray-500 h-full w-[30%]">
				<div id="action-buttons" hx-get="/action-buttons" hx-trigger="load" hx-swap="outerHTML" class="h-full w-full">
					<p>Loading action buttons...</p>
				</div>
			</div>

            <!-- Image (Placeholder) -->
			<div class="bg-blue-200 h-full w-[80%] text-white flex flex-row items-center justify-center">
				<div id="image-container" hx-get="/image" hx-trigger="load" hx-swap="outerHTML" class="h-[90%] w-[90%]">
					<p>Loading image...</p>
				</div>
			</div>

            <!-- Static HTML -->
            <div id="static-content" class="bg-gray-500 h-full w-[30%] text-white">
                <p>This is static content that loads immediately.</p>
            </div>
        </div>
    </body>
    </html>
    `
    w.Write([]byte(page))
}

func serveActionButtons(w http.ResponseWriter, r *http.Request) {
    // Simulate delay
    time.Sleep(1 * time.Second)

    buttons := `
    <div id="action-buttons" class="flex flex-col justify-evenly h-[100%] bg-gray-500 h-full w-[30%]">
        <button onclick="alert('Action 1')" class="bg-white w-fit">Action 1</button>
        <button onclick="alert('Action 2')" class="bg-white w-fit">Action 2</button>
    </div>
    `
    w.Write([]byte(buttons))
}

func serveImage(w http.ResponseWriter, r *http.Request) {
    // Simulate delay
    time.Sleep(4 * time.Second)

    image := `
    <div id="image-container" class="flex flex-row justify-center items-center h-[90%] w-[90%]">
        <img src="https://placehold.co/600x400" alt="Lazy Loaded Image" class="object-contain hover:object-scale-down">
    </div>
    `
    w.Write([]byte(image))
}
