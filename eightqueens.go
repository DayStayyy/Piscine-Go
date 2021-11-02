package piscine

import "github.com/01-edu/z01"

func EightQueens() {
	j := 1

	//Definition des tableaux
	var table = make([][]int, 8)
	for i := range table {
		table[i] = make([]int, 8)
	}

	for firstFor := 0; firstFor < 8; firstFor++ {
		table[0][firstFor] = 1
		j = 1
		for i := 1; i < len(table); i, j = i+1, j+1 {
			table[i][firstFor] = 2
			if !(firstFor+j >= 8) {
				table[i][firstFor+j] = 2
			}
			if firstFor-j >= 0 {
				table[i][firstFor-j] = 2
			}
		}

		for secondFor := 0; secondFor < 8; secondFor++ {
			if table[1][secondFor] != 2 {
				table[1][secondFor] = 1
				j = 1
				for i := 2; i < 8; i, j = i+1, j+1 {
					table[i][secondFor] = 2
					if secondFor+j < 8 {
						table[i][secondFor+j] = 2
					}
					if secondFor-j >= 0 {
						table[i][secondFor-j] = 2
					}
				}
				for thridFor := 0; thridFor < 8; thridFor++ {
					if table[2][thridFor] != 2 {
						table[2][thridFor] = 1
						j = 1
						for i := 3; i < 8; i, j = i+1, j+1 {
							table[i][thridFor] = 2
							if thridFor+j < 8 {
								table[i][thridFor+j] = 2
							}
							if thridFor-j >= 0 {
								table[i][thridFor-j] = 2
							}
						}
						for fourthFor := 0; fourthFor < 8; fourthFor++ {
							if table[3][fourthFor] != 2 {
								table[3][fourthFor] = 1
								j = 1
								for i := 4; i < 8; i, j = i+1, j+1 {
									table[i][fourthFor] = 2
									if fourthFor+j < 8 {
										table[i][fourthFor+j] = 2
									}
									if fourthFor-j >= 0 {
										table[i][fourthFor-j] = 2
									}
								}
								for fifthFor := 0; fifthFor < 8; fifthFor++ {
									if table[4][fifthFor] != 2 {
										table[4][fifthFor] = 1
										j = 1
										for i := 5; i < 8; i, j = i+1, j+1 {
											table[i][fifthFor] = 2
											if fifthFor+j < 8 {
												table[i][fifthFor+j] = 2
											}
											if fifthFor-j >= 0 {
												table[i][fifthFor-j] = 2
											}
										}
										for sixthFor := 0; sixthFor < 8; sixthFor++ {
											if table[5][sixthFor] == 0 {
												table[5][sixthFor] = 1
												j = 1
												for i := 6; i < 8; i, j = i+1, j+1 {
													table[i][sixthFor] = 2
													if sixthFor+j < 8 {
														table[i][sixthFor+j] = 2
													}
													if sixthFor-j >= 0 {
														table[i][sixthFor-j] = 2
													}
												}

												for seventhFor := 0; seventhFor < 8; seventhFor++ {
													if table[6][seventhFor] != 2 {
														table[6][seventhFor] = 1
														j = 1
														for i := 7; i < 8; i, j = i+1, j+1 {
															table[i][seventhFor] = 2
															if seventhFor+j < 8 {
																table[i][seventhFor+j] = 2
															}
															if seventhFor-j >= 0 {
																table[i][seventhFor-j] = 2
															}
														}

														for eighthFor := 0; eighthFor < 8; eighthFor++ {
															if table[7][eighthFor] == 0 {
																table[7][eighthFor] = 1
																z01.PrintRune(rune(firstFor + 49))
																z01.PrintRune(rune(secondFor + 49))
																z01.PrintRune(rune(thridFor + 49))
																z01.PrintRune(rune(fourthFor + 49))
																z01.PrintRune(rune(fifthFor + 49))
																z01.PrintRune(rune(sixthFor + 49))
																z01.PrintRune(rune(seventhFor + 49))
																z01.PrintRune(rune(eighthFor + 49))
																z01.PrintRune('\n')
															}

														}
														//si ca marche pas 7
														j = 1
														for i := 7; i < 8; i, j = i+1, j+1 {
															table[i][sixthFor] = 0
															if seventhFor+j < 8 {
																table[i][seventhFor+j] = 0
															}
															if seventhFor-j >= 0 {
																table[i][seventhFor-j] = 0
															}
														}
														j = 1
														for i := 1; i < len(table); i, j = i+1, j+1 {
															table[i][firstFor] = 2
															if !(firstFor+j >= 8) {
																table[i][firstFor+j] = 2
															}
															if firstFor-j >= 0 {
																table[i][firstFor-j] = 2
															}
														}
														j = 1
														for i := 2; i < 8; i, j = i+1, j+1 {
															table[i][secondFor] = 2
															if secondFor+j < 8 {
																table[i][secondFor+j] = 2
															}
															if secondFor-j >= 0 {
																table[i][secondFor-j] = 2
															}
														}
														j = 1
														for i := 3; i < 8; i, j = i+1, j+1 {
															table[i][thridFor] = 2
															if thridFor+j < 8 {
																table[i][thridFor+j] = 2
															}
															if thridFor-j >= 0 {
																table[i][thridFor-j] = 2
															}
														}
														j = 1
														for i := 4; i < 8; i, j = i+1, j+1 {
															table[i][fourthFor] = 2
															if fourthFor+j < 8 {
																table[i][fourthFor+j] = 2
															}
															if fourthFor-j >= 0 {
																table[i][fourthFor-j] = 2
															}
														}
														for i := 5; i < 8; i, j = i+1, j+1 {
															table[i][fifthFor] = 2
															if fifthFor+j < 8 {
																table[i][fifthFor+j] = 2
															}
															if fifthFor-j >= 0 {
																table[i][fifthFor-j] = 2
															}
														}
														j = 1
														for i := 6; i < 8; i, j = i+1, j+1 {
															table[i][sixthFor] = 2
															if sixthFor+j < 8 {
																table[i][sixthFor+j] = 2
															}
															if sixthFor-j >= 0 {
																table[i][sixthFor-j] = 2
															}
														}
													}
												}
												//si ca marche pas 6
												j = 1
												for i := 6; i < 8; i, j = i+1, j+1 {
													table[i][sixthFor] = 0
													if sixthFor+j < 8 {
														table[i][sixthFor+j] = 0
													}
													if sixthFor-j >= 0 {
														table[i][sixthFor-j] = 0
													}
												}
												j = 1
												for i := 1; i < len(table); i, j = i+1, j+1 {
													table[i][firstFor] = 2
													if !(firstFor+j >= 8) {
														table[i][firstFor+j] = 2
													}
													if firstFor-j >= 0 {
														table[i][firstFor-j] = 2
													}
												}
												j = 1
												for i := 2; i < 8; i, j = i+1, j+1 {
													table[i][secondFor] = 2
													if secondFor+j < 8 {
														table[i][secondFor+j] = 2
													}
													if secondFor-j >= 0 {
														table[i][secondFor-j] = 2
													}
												}
												j = 1
												for i := 3; i < 8; i, j = i+1, j+1 {
													table[i][thridFor] = 2
													if thridFor+j < 8 {
														table[i][thridFor+j] = 2
													}
													if thridFor-j >= 0 {
														table[i][thridFor-j] = 2
													}
												}
												j = 1
												for i := 4; i < 8; i, j = i+1, j+1 {
													table[i][fourthFor] = 2
													if fourthFor+j < 8 {
														table[i][fourthFor+j] = 2
													}
													if fourthFor-j >= 0 {
														table[i][fourthFor-j] = 2
													}
												}
												for i := 5; i < 8; i, j = i+1, j+1 {
													table[i][fifthFor] = 2
													if fifthFor+j < 8 {
														table[i][fifthFor+j] = 2
													}
													if fifthFor-j >= 0 {
														table[i][fifthFor-j] = 2
													}
												}
											}
										}

										//si ca marche pas 5
										j = 1
										for i := 5; i < 8; i, j = i+1, j+1 {
											table[i][fifthFor] = 0
											if fifthFor+j < 8 {
												table[i][fifthFor+j] = 0
											}
											if fifthFor-j >= 0 {
												table[i][fifthFor-j] = 0
											}
										}
										j = 1
										for i := 1; i < len(table); i, j = i+1, j+1 {
											table[i][firstFor] = 2
											if !(firstFor+j >= 8) {
												table[i][firstFor+j] = 2
											}
											if firstFor-j >= 0 {
												table[i][firstFor-j] = 2
											}
										}
										j = 1
										for i := 2; i < 8; i, j = i+1, j+1 {
											table[i][secondFor] = 2
											if secondFor+j < 8 {
												table[i][secondFor+j] = 2
											}
											if secondFor-j >= 0 {
												table[i][secondFor-j] = 2
											}
										}
										j = 1
										for i := 3; i < 8; i, j = i+1, j+1 {
											table[i][thridFor] = 2
											if thridFor+j < 8 {
												table[i][thridFor+j] = 2
											}
											if thridFor-j >= 0 {
												table[i][thridFor-j] = 2
											}
										}
										j = 1
										for i := 4; i < 8; i, j = i+1, j+1 {
											table[i][fourthFor] = 2
											if fourthFor+j < 8 {
												table[i][fourthFor+j] = 2
											}
											if fourthFor-j >= 0 {
												table[i][fourthFor-j] = 2
											}
										}
									}
								}

								//si ca marche pas 4
								j = 1
								for i := 4; i < 8; i, j = i+1, j+1 {
									table[i][fourthFor] = 0
									if fourthFor+j < 8 {
										table[i][fourthFor+j] = 0
									}
									if fourthFor-j >= 0 {
										table[i][fourthFor-j] = 0
									}
								}
								j = 1
								for i := 1; i < len(table); i, j = i+1, j+1 {
									table[i][firstFor] = 2
									if !(firstFor+j >= 8) {
										table[i][firstFor+j] = 2
									}
									if firstFor-j >= 0 {
										table[i][firstFor-j] = 2
									}
								}
								j = 1
								for i := 2; i < 8; i, j = i+1, j+1 {
									table[i][secondFor] = 2
									if secondFor+j < 8 {
										table[i][secondFor+j] = 2
									}
									if secondFor-j >= 0 {
										table[i][secondFor-j] = 2
									}
								}
								j = 1
								for i := 3; i < 8; i, j = i+1, j+1 {
									table[i][thridFor] = 2
									if thridFor+j < 8 {
										table[i][thridFor+j] = 2
									}
									if thridFor-j >= 0 {
										table[i][thridFor-j] = 2
									}
								}
							}
						}

						//si ca marche pas 3
						j = 1
						for i := 3; i < 8; i, j = i+1, j+1 {
							table[i][thridFor] = 0
							if thridFor+j < 8 {
								table[i][thridFor+j] = 0
							}
							if thridFor-j >= 0 {
								table[i][thridFor-j] = 0
							}
						}
						j = 1
						for i := 1; i < len(table); i, j = i+1, j+1 {
							table[i][firstFor] = 2
							if !(firstFor+j >= 8) {
								table[i][firstFor+j] = 2
							}
							if firstFor-j >= 0 {
								table[i][firstFor-j] = 2
							}
						}
						j = 1
						for i := 2; i < 8; i, j = i+1, j+1 {
							table[i][secondFor] = 2
							if secondFor+j < 8 {
								table[i][secondFor+j] = 2
							}
							if secondFor-j >= 0 {
								table[i][secondFor-j] = 2
							}
						}
					}
				}

				//si ca marche pas 2
				j = 1
				for i := 2; i < 8; i, j = i+1, j+1 {
					table[i][secondFor] = 0
					if secondFor+j < 8 {
						table[i][secondFor+j] = 0
					}
					if secondFor-j >= 0 {
						table[i][secondFor-j] = 0
					}
				}
				j = 1
				for i := 1; i < len(table); i, j = i+1, j+1 {
					table[i][firstFor] = 2
					if !(firstFor+j >= 8) {
						table[i][firstFor+j] = 2
					}
					if firstFor-j >= 0 {
						table[i][firstFor-j] = 2
					}
				}
			}
		}

		//si ca marche pas 1
		j = 1
		for i := 1; i < 8; i, j = i+1, j+1 {
			table[i][firstFor] = 0
			if firstFor+j < 8 {
				table[i][firstFor+j] = 0
			}
			if firstFor-j >= 0 {
				table[i][firstFor-j] = 0
			}
		}
	}
}
