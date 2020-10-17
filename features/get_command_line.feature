Feature: Get command line arguments
  Make sure command line parameters are validated and used

  Scenario: No parameters
    Given no parameters
    When I count my parameters
    Then I should see 0
