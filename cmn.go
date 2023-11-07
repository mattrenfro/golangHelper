package cmn
 import(
   "strings"
   "fmt"
 	"log"
 	"os"
  "os/exec"
 	"runtime"
 	"syscall"
  "path/filepath"
  "io"
  "time"
 )

 func Chompr(text string) string{
   text = strings.TrimSuffix(text, "\n")
   text = strings.TrimSuffix(text, "\r")
   return text
 }

 func Countdown(seconds int){
   fmt.Print("\n")
   //counter := 0
  for i := seconds; i > 0; i-- {
    //counter++
    spaces := strings.Repeat("-", i)

    fmt.Print("\r")
    fmt.Print(spaces)
    time.Sleep(time.Second/2)
    fmt.Print("\r")
    spaces = strings.Repeat(" ", i)
    fmt.Print(spaces)
  }
 }

 func TerminalClear(){
   if runtime.GOOS == "windows" {
     cmd := exec.Command("cmd", "/c", "cls")
     cmd.Stdout = os.Stdout
     cmd.Run()
   }else{
     fmt.Println("\\033c")
   }
 }

 func SetWritable(filepath string) error {
 err := os.Chmod(filepath, 0222)
 return err
}

func SetReadOnly(filepath string) error {
 err := os.Chmod(filepath, 0444)
 return err
}

func IsHidden(filename string) (bool, error) {

	if runtime.GOOS == "windows" {
 		pointer, err := syscall.UTF16PtrFromString(filename)
 		if err != nil {
 			return false, err
 		}
 		attributes, err := syscall.GetFileAttributes(pointer)
 		if err != nil {
 			return false, err
 		}
 		return attributes&syscall.FILE_ATTRIBUTE_HIDDEN != 0, nil
 	} else {
 		// unix/linux file or directory that starts with . is hidden
 		if filename[0:1] == "." {
 			return true, nil

 		} else {
 			return false, nil
 		}

 	}
  return false, nil
 }

 func HideFile(filename string) error {
   if runtime.GOOS == "windows" {
    filenameW, err := syscall.UTF16PtrFromString(filename)
    if err != nil {
        return err
    }
    err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_HIDDEN)
    if err != nil {
        return err
    }
  }else{
    if !strings.HasPrefix(filepath.Base(filename), ".") {
        HideFile(filename)
    }
  }
  return nil
}

// if os is windows, set the file to FILE_ATTRIBUTE_NORMAL
// if linux and you want to turn off read only, then pass in true for 2nd arg
// if linux and you want to UnHideFile, then pass in false for 2nd arg
func RevealFile(filename string, readable bool) error {
  if runtime.GOOS == "windows" {
   filenameW, err := syscall.UTF16PtrFromString(filename)
   if err != nil {
       return err
   }
   err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_NORMAL)
   if err != nil {
       return err
   }
 }else{
   if(!readable){
     if strings.HasPrefix(filepath.Base(filename), ".") {
         UnHideFile(filename)
     }
  }else{
    SetWritable(filepath.Base(filename))
  }
 }
 return nil
}

func MakeReadOnly(filename string) error {
  if runtime.GOOS == "windows" {
   filenameW, err := syscall.UTF16PtrFromString(filename)
   if err != nil {
       return err
   }
   err = syscall.SetFileAttributes(filenameW, syscall.FILE_ATTRIBUTE_READONLY)
   if err != nil {
       return err
   }
 }else{
   SetReadOnly(filepath.Base(filename))
 }
 return nil
}

 func CanWrite(filepath string) (bool, error) {
 	file, err := os.OpenFile(filepath, os.O_WRONLY, 0666)
 	if err != nil {
 		if os.IsPermission(err) {
 			return false, err
 		}
 	}
 	file.Close()
 	return true, nil
 }

 /*func SetWritable(filepath string) error {
 	err := os.Chmod(filepath, 0222)
 	return err
 }

 func SetReadOnly(filepath string) error {
 	err := os.Chmod(filepath, 0444)
 	return err
 }*/

 func MakeFileHidden(filename string){
  newName := "." + filename
 	err := os.Rename(filename, newName)
 	if err != nil {
 		log.Fatal(err)
 	}
 }

 func UnHideFile(filename string){
	newName := strings.TrimLeft(filename, ".")
  fmt.Println(newName)
	err := os.Rename(filename, newName)
	if err != nil {
		log.Fatal(err)
	}
 }

 // This function tries to replicate the shift function in Perl by removing the first
// element of a slice and returning it
func Shift (pToSlice *[]string) string {
    sValue := (*pToSlice)[0]

    return sValue
}

func SubstringLeft(s string, i int) string{
  return s[:i]
}

func SubstringRight(s string, i int) string{
  return s[i:]
}

func SubstringInfix(s string, i int, j int) string{
  return s[i:j]
}

func CutSlice(num int, a []string) []string{
  return a[num:min(num, len(a))]
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func WriteToFile(filename string, data string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = io.WriteString(file, data)
    if err != nil {
        return err
    }
    return file.Sync()
}
