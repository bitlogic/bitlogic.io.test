Feature: Landing Page
  In order to communicate who we are and to promote our services to potential customers
  i need to be able to navigate the bitlogic landing Page

  Scenario: Show Home Page 
    Given i'm interested in bitlogic.io 
    When i navigate to http://bitlogic.io 
    Then the landing page is loaded
    