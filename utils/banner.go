package utils

const author string = "[+]Author:devpwn\n"
const version string = "[+]Version: 2.0\n\n"

// UglyBanner banner
func UglyBanner() string {

	var banner string = `
	

██╗  ██╗███╗   ███╗██╗     ██████╗ ██████╗  ██████╗ ███████╗ ██████╗ █████╗ ███╗   ██╗
╚██╗██╔╝████╗ ████║██║     ██╔══██╗██╔══██╗██╔════╝ ██╔════╝██╔════╝██╔══██╗████╗  ██║
 ╚███╔╝ ██╔████╔██║██║     ██████╔╝██████╔╝██║█████╗███████╗██║     ███████║██╔██╗ ██║
 ██╔██╗ ██║╚██╔╝██║██║     ██╔══██╗██╔═══╝ ██║╚════╝╚════██║██║     ██╔══██║██║╚██╗██║
██╔╝ ██╗██║ ╚═╝ ██║███████╗██║  ██║██║     ╚██████╗ ███████║╚██████╗██║  ██║██║ ╚████║
╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝╚═╝  ╚═╝╚═╝      ╚═════╝ ╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝                                                                       	


`
	returnBanner := banner + author + version
	return returnBanner
}
