package batteryman

import (
	"fmt"
	"io/fs"
	"os"
	"strconv"
	"strings"
)

// Noteworthy mention: All of the amazing wiki maintainers at the arch wiki.
// https://wiki.archlinux.org/title/Laptop/ASUS#ZenBook

// The arch wiki does state that these changes were brought in Kernel 5.4, so if you are running a LTS distro
// you may face some issues. Run "uname -r" to check compatibility.

// Rest assured, if you are running a more cutting edge distro such as Arch, or Fedora, compatibility is guaranteed.

func GetBatteryDevice() (string, error) {
	// Purpose of this function:
	// Fetch battery device name, whether its BAT0, BAT1, BATC or BATT.
	// Returns: Battery device name, or an error.

	// List of all our potential devices.
	devices := []string{"BAT0", "BAT1", "BATC", "BATT"}

	// Iterate through all battery devices.
	for _, device := range devices {
		// Read the status of said battery device.
		_, err := os.ReadFile(fmt.Sprintf("/sys/class/power_supply/%s/status", device))
		if err != nil {
			// This battery device does not exist.
			// Skip to the next iteration.
			continue
		}

		// Successful battery device found. Return.
		return device, nil
	}

	// No battery device was found.
	return "", fmt.Errorf("battery device was unable to be found")
}

func GetThreshold() (uint, error) {
	// Purpose of this function:
	// To allow you to read the current battery charge limit.

	// Fetch battery device.
	device, err := GetBatteryDevice()
	if err != nil {
		return 0, err
	}

	// Read from said system location.
	contents, err := os.ReadFile(fmt.Sprintf("/sys/class/power_supply/%s/charge_control_end_threshold", device))
	if err != nil {
		return 0, err
	}

	// This will produce newline characters at the end.
	// We must trim any whitespace, so Atoi can successfully convert our string into an integer.
	sanitised := strings.TrimSpace(string(contents))

	// Attempt to convert the string threshold into an integer.
	threshold, err := strconv.Atoi(sanitised)
	if err != nil {
		return 0, err
	}

	return uint(threshold), nil
}

func SetThreshold(threshold uint) error {
	// Purpose of this function:
	// To allow you to set a battery charge limit.

	// If a threshold below 1 is set, it may produce very funky behaviour.
	// As laptops do not have a way of correctly determining charge, we will set the lower limit to 5%.
	// The upper limit will be 100%.
	if threshold < 5 || threshold > 100 {
		return fmt.Errorf("threshold may not preceed 1 or exceed 100")
	}

	// Fetch battery device.
	device, err := GetBatteryDevice()
	if err != nil {
		return err
	}

	// Convert the threshold into a string, and write it to our system location.
	// This process will likely throw an error if superuser privileges are unavailable.
	result := strconv.Itoa(int(threshold))
	err = os.WriteFile(fmt.Sprintf("/sys/class/power_supply/%s/charge_control_end_threshold", device), []byte(result), fs.ModePerm)
	if err != nil {
		return err
	}

	// No error produced. Continue.
	return nil
}

func GetStatus() (string, error) {
	// Purpose of this function:
	// Allows you to check the system charge status.

	// Fetch battery device.
	device, err := GetBatteryDevice()
	if err != nil {
		return "", err
	}

	// Read from said system location.
	contents, err := os.ReadFile(fmt.Sprintf("/sys/class/power_supply/%s/status", device))
	if err != nil {
		return "", err
	}

	// This will produce newline characters at the end.
	// Trim whitespace for convenience reasons.
	sanitised := strings.TrimSpace(string(contents))

	// Return our sanitised output.
	return sanitised, nil
}

func GetCapacity() (uint, error) {
	// Purpose of this function:
	// Check your current charge levels.

	// cat /sys/class/power_supply/BAT0/capacity

	// Fetch battery device.
	device, err := GetBatteryDevice()
	if err != nil {
		return 0, err
	}

	// Read from said system location.
	contents, err := os.ReadFile(fmt.Sprintf("/sys/class/power_supply/%s/capacity", device))
	if err != nil {
		return 0, err
	}

	// This will produce newline characters at the end.
	// Trim whitespace for convenience reasons.
	sanitised := strings.TrimSpace(string(contents))

	// Attempt to convert the string capacity into an integer.
	capacity, err := strconv.Atoi(sanitised)
	if err != nil {
		return 0, err
	}

	return uint(capacity), nil
}
