// DataHandler.js
export class DataHandler {
    constructor(data, options) {
      this.data = data;
      this.rowsPerPage = options.rowsPerPage || 10;
      this.currentPage = 1;
    }
  
    getRows() {
      const startIdx = (this.currentPage - 1) * this.rowsPerPage;
      const endIdx = startIdx + this.rowsPerPage;
      return this.data.slice(startIdx, endIdx);
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
  