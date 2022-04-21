describe("renders the bookmarks page", () => {
    it("renders correctly", () => {
      cy.visit("/bookmarks");
      cy.contains("bookmark").should("exist");
      
      });

      
});