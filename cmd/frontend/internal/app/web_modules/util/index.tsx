/**
 * supportedExtensions are the file extensions
 * the extension will apply annotations to
 */
export const supportedExtensions = new Set<string>([
	"go", // Golang
	"ts", "tsx", // TypeScript
	"js", "jsx", // JavaScript
	"java", // Java
	"py", // Python
	"php", // PHP
]);

/**
 * getModeFromExtension returns the LSP mode for the
 * provided file extension (e.g. "jsx")
 */
export function getModeFromExtension(ext: string): string {
	switch (ext) {
		case "go":
			return "go";
		case "ts":
		case "tsx":
			return "typescript";
		case "js":
		case "jsx":
			return "javascript";
		case "java":
			return "java";
		case "py":
		case "pyc":
		case "pyd":
		case "pyo":
		case "pyw":
		case "pyz":
			return "python";
		case "php":
		case "phtml":
		case "php3":
		case "php4":
		case "php5":
		case "php6":
		case "php7":
		case "phps":
			return "php";
		default:
			return "unknown";
	}
}

export function getPathExtension(path: string): string {
	const pathSplit = path.split(".");
	if (pathSplit.length === 1) {
		return "";
	}
	if (pathSplit.length === 2 && pathSplit[0] === "") {
		return ""; // e.g. .gitignore
	}
	return pathSplit[pathSplit.length - 1].toLowerCase();
}
