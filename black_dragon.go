package main

/*
	Things to Work On:
		1. How to make EXE out of this
 */

/*
	The main libraries used in Black Dragon. Might change due to a possible dependency error in the future.
	All of these libraries as of 1/10/2022 are not permanent.
*/
import (
  "bufio"
  "fmt"
  "os"
  "log"
  "os/user"
  "strings"
  "time"
  "io/ioutil"
  "net"
  "net/http"
  "encoding/json"
  "github.com/shirou/gopsutil/cpu"
  "github.com/shirou/gopsutil/mem"
  "github.com/shirou/gopsutil/host"
  "github.com/shirou/gopsutil/disk"
  "crypto"
  "crypto/rsa"
  "crypto/rand"
  "crypto/sha256"
)

// Create two global variables used for encrypt/decrypt
var key *rsa.PrivateKey = nil
var bytes []byte = nil

/*
	Takes the error and outputs the given error if there is any.
 */
func fatal_err(err error){
	// check if error is not null to send error back
	if err != nil {
		log.Fatalf(err.Error())
	}
}

/* 
	The main startup screen that appears upon launching Black Dragon. 
 */
func startup(){
	// get the user from the user library (goes according to the user of the system)
	user, err := user.Current()
	fatal_err(err)

	// Main log in screen that pops up upon every log in
	fmt.Println("--IMPORTANT:-------------------------------------------------")
	fmt.Println("Welcome to Black Dragon")
	fmt.Println("The Ultimate Computer Networking tool at your finger tips.")
	println("                                        *******eeee")
	println("                                          ***$$$$$$$$e")
	println("                                                **$$$$$e")
	println("           e$$$e                                   *$$$$$")
	println("          $$$$                                     e$$$$$")
	println("         $$$$                                  ee$$$$$$$")
	println("         $$$$               ee***ee***eeeee$$$$$$$$$$$*")
	println("         $$$$e        eee$$$     $     $$$$$$$$$$$$$*        ee")
	println("         *$$$$$$$$$$$$$$$$$$ $$$ $ $$$ $$$$$$$$**      ee****WW$")
	println(" eeeeee    *$$$$$$$$$$$$$$$$  $$ $  $$ $*****        e*!!!W**!W*")
	println("$!WWWW!***ee   *************eee$$**eee$             e*!W**!!!!$")
	println(" $!!!!***WW!**eee        e**!!!!!!!!!!!***e      e**!!W*!!!!!!$")
	println(" $!!!!!!!!!*W!!!!$      $!!!!W****WWWW!!!!!$     *W!!!*WW!!!!!$")
	println(" $!!!!!!!!!WW*!!!$      *W!!*W!!!!!!!!*W!!W*     e*!!!!!!*W!!!$")
	println(" $!!!!!W***!!WW!!*e       ******WWWWW*****       $!!!W***W!**W*")
	println(" $!!!W*WW****!!*W!!*eee    ee   $!!$   ee     ee*!!W*!!!!!*W!!$")
	println("  $W*!$!!!!!!!!!!*WW!!!****!W***!!!!***W!*****!!!!$!!!!!!!!*W$")
	println("  $!W*!!!!!!!!!!!W**!WW!!!!W*!!!!!!!!!!*WW*****W!!$!!!!!!!!!W*")
	println("   *WW!!!!!!!!!W*!W**!!***W$!!!!W!!!W!!!$!!!!!!$!W*!!!!!!W**")
	println("      ****WW!W*!W$WWW!!!!!!$!!!!*WWW*!!!$!!!!!!$W$!!!WWWW*")
	println("            *$W*!!!!!*WW!!!$!!!!!!$!!!!W*!!WW**!!!*$*")
	println("              $!!!!!!!!!*W!*W!!!!!$!!!!$!W*!!!!!!!!!$")
	println("              *W!!!!!!!!!!***W!!!!$!!!!$*!!!!!!!!!!!$")
	println("               $!!!!!!!!!!!!!$!!!!$!!!W*!!!!!!!!!!!$")
	println("                $!!!!!*W!!!!!$!!!!$!!!$!!!!!W!!!!!$")
	println("eeeeeeee         $!!!!!!$!!!!$!!!!$!!!$!!!!W*!!!!W*")
	println("$!!!!W*       ee$W$!!!!!!$!!!$!!!!$!!!$!!!W*!!!!W*")
	println("$!!!*ee**e**e*WW*!!$!!!!!!$!!$!!!!$!!!$!!!$!!!WW*")
	println("$!$WW!!*******!!WWW**W*W!!$WW*!!!!$!!!$!W$WW**W!*e")
	println("$*   **WWWWWW***e*!!$!!$**!W*!!!WW*$!!!*W!!*W!!*W!*e")
	println("              e*!!!$!!W*!W*!!!!$    $!!!!$!!!$!!!$!*e")
	println("              *WWW*WW*WWW*WWWW*      *WWW*WWW*WWW*WW*")
	println("Credit: boba@gagme.wwa.com\n")
	fmt.Println("Enjoy the power >:)")
	fmt.Println("--------------------------------------------------------------")
	fmt.Println("\n\n\nUSER is: @", user.Username) // use the user's username from their machine to display in the front screen
	fmt.Print("\n")
	

	// Delay before getting rid of the login screen
	time.Sleep(2*time.Second) // have a delay 
	fmt.Print("\033[H\033[2J")
}

/*
	Checks if a given website is reachable and gives its corresponding HTTP status Code.
	Websites are given in the format of https://www.<website>.<domain>
 */
func checkIfReachable(website string) int {
	// get the response from the http library
	response, errors := http.Get(website)
	fatal_err(errors)

	// check the status code to return the respective message
	if response.StatusCode == 200{
		println("Requested Website is Running Properly")
		return 0
	} else if response.StatusCode == 301 {
		println("Appears the website you requested has permanently been redirected")
		return 0
	} else if response.StatusCode == 302 {
		println("Appears the website you are looking for has been temporarily redirected")
		return 0
	} else if response.StatusCode == 404 {
		println("Appears the website you are looking for doesn't exist")
		return 0
	} else if response.StatusCode == 410 {
		println("Appears the website has disappeared from the internet :(")
		return 0
	} else if response.StatusCode == 500 {
		println("Oh no Internal Server Error, check if you are connected to the internet")
		return 0
	} else if response.StatusCode == 503 {
		println("This service is not available because you are not connected to the internet, try again after you have connected to the internet")
		return 0
	}

	// return 0 at the end for success
	return 0
}

/*
	Processes the given input and outputs the requested information. Returns 0 at the end.
	Input is given in the format command arg1 arg2 arg3 arg4 ...
	** NEED TO WORK ON GIVEN ERROR OUTPUTS **
 */
func process_input(input string) int{
	
	// First handle single inputs 

	// clear screen command
	if strings.Compare("cls", input) == 0 || strings.Compare("clear", input) == 0 {
		fmt.Print("\033[H\033[2J")
	}
	
	// List files/subdirectories command
	if strings.Compare("ls", input) == 0 {
		files, err := ioutil.ReadDir("./")
		fatal_err(err)
		for _, f := range files {
            fmt.Println(f.Name())
   		}
	}


	// List directory command 
	if strings.Compare("pwd", input) == 0 {
		dirname, err := os.Getwd()
		fatal_err(err)
		fmt.Println(dirname)
	}

	// key generation command
	if strings.Compare("generate", input) == 0 {
		// Generate a private key from the RSA library
		priKey, err := rsa.GenerateKey(rand.Reader, 2048)
		fatal_err(err)

		key = priKey // store the private key in the global instance to use

		// Use the private key struct to generate a public key
		pubKey := priKey.PublicKey

		// print so the user can use it
		fmt.Println("Public Key: %s\n", pubKey)
		fmt.Println("Private Key: %s\n", priKey)
	}

	// Now we get rid of the single inputs, lets deal with multiple arguments (still simple commands)
	// To begin lets break up the string first into [command, arg1, arg2, arg3, ...]
	var new_input = strings.Split(input, " ")
	
	// Simple multiple arguments commands
	if strings.Compare("-h", new_input[0]) == 0 || strings.Compare("help", new_input[0]) == 0{
		// print all the commands
		println("List of Commands in Black Dragon:\n")
		println("Basic Commands:")
		println("cls - clear the screen")
		println("clear - clear the screen")
		println("ls - list all the files and directories in your current directory")
		println("pwd - print the current directory")
		println("cd arg1 - change directory to arg1")
		println("rm -f arg1 - remove choosen file")
		println("rm -d arg1 - remove choosen directory")
		println("touch arg1 - creates a new requested file")
		println("\n")
		println("Networking Commands:")
		println("check arg1 - checks the http status of a requested website")
		println("ping arg1 - checks if a ping can be made from the machine to the requested server")
		println("\n")
		println("Security Commands")
		println("generate - generate a private key and a public key for encryption/decryption")
		println("code arg1 - encrypts the password argument given")
		println("crack arg1 - decrypts the encryption key given")
		println("\n")
	}

	// Make directory command
	if strings.Compare("mkdir", new_input[0])  == 0 {
		os.Mkdir(new_input[1], 0755) // Direct command from OS library
	}

	// Remove command
	if strings.Compare("rm", new_input[0]) == 0 {
		// depending on the flag either remove directory or files
		if strings.Compare("-d", new_input[1]) == 0 {
			os.RemoveAll(new_input[2]) // RemoveAll because its a directory
		} else if strings.Compare("-f", new_input[1]) == 0 {
			os.Remove(new_input[2]) // Remove just removes the single file
		}
	}

	// Change directory command
	if strings.Compare("cd", new_input[0]) == 0 {
		os.Chdir(new_input[1]) // Changes the directory from the one in the main function
	}

	// touch command
	if strings.Compare("touch", new_input[0]) == 0 {
		f, e := os.Create(new_input[1])
		fatal_err(e)
		defer f.Close()
	}

	// More Advance multiple argument commands

	// HTTP status code command
	if strings.Compare("check", new_input[0]) == 0 {
		checkIfReachable(new_input[1])
	}

	// Ping command
	if strings.Compare("ping", new_input[0]) == 0 {
		// Get the host name
		hostName := new_input[1]
		timeOut := time.Duration(5) * time.Second
		
		// Create a tcp connection to test
		conn, err := net.DialTimeout("tcp", hostName+":80", timeOut)
		fatal_err(err)
		
		// Print out the results of the ping
		fmt.Printf("Connection established between %s and localhost with time out of %d seconds.\n", hostName, int64(5))
		fmt.Printf("Home Address: %s \n", conn.LocalAddr().String())
		fmt.Printf("Requested Address: %s \n", conn.RemoteAddr().String())
	}

	// System information command
	if strings.Compare("sysinfo", new_input[0]) == 0 {
		// print out cpu information
		if strings.Compare("cpu", new_input[1]) == 0 {
			cpuInfos, err := cpu.Info()
			fatal_err(err)
			for _, ci := range cpuInfos {
				s, _ := json.MarshalIndent(ci, "", "\t")
    			println(string(s))
			}
		// print out memory information
		} else if strings.Compare("mem", new_input[1]) == 0 {
			memInfo, _ := mem.VirtualMemory()
			s, _ := json.MarshalIndent(memInfo, "", "\t")
    		println(string(s))
		// print out host address information
		} else if strings.Compare("host", new_input[1]) == 0 {
			hInfo, _ := host.Info()
			s, _ := json.MarshalIndent(hInfo, "", "\t")
    		println(string(s))
		// print out disk/storage information
		} else if strings.Compare("disk", new_input[1]) == 0 {
			parts, err := disk.Partitions(true)
			fatal_err(err)
			for _, part := range parts {
				fmt.Printf("part:%v\n", part.String())
				diskInfo, _ := disk.Usage(part.Mountpoint)
				s, _ := json.MarshalIndent(diskInfo, "", "\t")
    			println(string(s))
			}

			ioStat, _ := disk.IOCounters()
			for k, v := range ioStat {
				fmt.Printf("%v:%v\n", k, v)
			}
		}
	}

	// encrypt command
	if strings.Compare("code", new_input[0]) == 0 {
		// using the global instance of the private key encrypt the string
		encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &key.PublicKey, []byte(new_input[1]), nil)
		fatal_err(err)

		bytes = encryptedBytes // store the encrypted bytes into bytes global instance
		
		// print the encrypted bytes for the user
		fmt.Println("encrypted bytes: ", encryptedBytes)	
	} 

	// decrypt command
	if strings.Compare("crack", new_input[0]) == 0 {
		// using the global instance of the private key decrypt the encryptedbytes
		decryptedBytes, err := key.Decrypt(nil, bytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
		fatal_err(err)

		// print out the decrypted message
		fmt.Println("decrypted message: ", string(decryptedBytes))
	}

	return 0 // return 0 upon successful or even failure, fatal_err function will return failure for you
}

/*
	Main function that launches Black Dragon
 */
func main() {

  // Create a reading function to read any input given by the user
  reader := bufio.NewReader(os.Stdin)
  startup()

  // Create an infinite loop to keep user input until the program is exited out of
  for {
	// Get the directory name to print out
	dirname, err := os.Getwd()
  	fatal_err(err)
    fmt.Print(dirname+">")
    
	// Take the input given thru stdin and replace anything unnecessary
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	
	// Process the input and give the user the requested information
	process_input(input[:len(input)-1])

  }

}