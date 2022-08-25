describe('empty spec', () => {
    it('passes', () => {
      cy.visit('http://localhost:3000/')
      cy.get('.payments-link').first().click()
      cy.url().should('include', '/payments')
    })
  })