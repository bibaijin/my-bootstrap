filename=An-Experimental-System-for-Generating-and-Measuring-Tunable-Orbital-Angular-Momentum-in-Radio

main.pdf: main.tex ref.bib
	latexmk -xelatex main.tex

save:
	cp main.pdf "$(filename).pdf"

view:
	evince main.pdf &

clean:
	rm -f *.aux *.bbl *.blg *.idx *.ilg *.ind *.out *.toc *.nav *.snm main.pdf *.run.xml *-blx.bib *.bcf *.log *.fdb_latexmk *.fls
