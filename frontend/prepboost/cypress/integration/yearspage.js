describe("renders the home page", () => {
    it("renders correctly", () => {
      cy.visit("/years");
      cy.contains("different years").should("exist");
      cy.findAllByText("2021").should("exist");
      cy.findAllByText("2020").should("exist");
      cy.findAllByText("2019").should("exist");
      cy.findAllByText("2011").should("exist");
      });

      it("routes to a correct pages", () => {
           cy.findAllByText("2021").click();
           cy.url().should("include", "topicquestions");
         });
});