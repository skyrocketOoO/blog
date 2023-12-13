import { writable } from 'svelte/store';

export class DataHandler {
  constructor(data, options) {
    this.data = data;
    this.rowsPerPage = options.rowsPerPage || 10;
    this.currentPage = 1;
  }

  getRows() {
    const startIdx = (this.currentPage - 1) * this.rowsPerPage;
    const endIdx = startIdx + this.rowsPerPage;
    const slicedRows = this.data.slice(startIdx, endIdx);
    rows.set(slicedRows); // Set the store value with sliced rows
    return slicedRows; // Return sliced rows
  }

  nextPage() {
    const totalPages = Math.ceil(this.data.length / this.rowsPerPage);
    if (this.currentPage < totalPages) {
      this.currentPage++;
    }
  }

  prevPage() {
    if (this.currentPage > 1) {
      this.currentPage--;
    }
  }
}


export const rows = writable([]);