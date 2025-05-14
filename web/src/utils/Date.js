import moment from "moment";

export const formatDateTime = (datetime) => {
  return moment(datetime).format("DD-MMMM-YYYY"); // Example: January 1, 2025 3:30 PM
};


export const formatTime = (datetime) => {
    return moment(datetime).format("dddd, h:mm A")
}