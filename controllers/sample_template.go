package controllers

import (
	"log"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

func HTMLtoPDF(c *gin.Context) {
	wkhtmltopdf.SetPath("C:\\Program Files\\wkhtmltopdf\\bin\\wkhtmltopdf.exe")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	htmlContent := `
<html>
<head>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f4f4f4;
        }
        .container {
            display: flex;
            flex-direction: column;
            width: 80%;
            margin: 30px auto;
            background-color: white;
            border-radius: 10px;
            box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
            padding: 20px;
        }
        .header {
            text-align: center;
            font-size: 24px;
            color: #4CAF50;
            margin-bottom: 20px;
        }

        /* Flexbox for Employee Information */
        .employee-info {
            display: flex;
            flex-wrap: wrap;
            justify-content: space-between;
            gap: 20px;
            margin-bottom: 20px;
        }
        .employee-info div {
            flex: 1 1 45%;
        }

        .section h3 {
            font-size: 20px;
            color: #333;
            margin-bottom: 10px;
        }

        /* CSS Grid for Tables */
        .table {
            display: grid;
            grid-template-columns: 1fr 1fr;
            grid-gap: 10px;
            width: 100%;
            border-collapse: collapse;
            margin-bottom: 20px;
        }
        .table div {
            padding: 8px;
            border: 1px solid #ddd;
        }
        .table .header {
            background-color: #f2f2f2;
            font-weight: bold;
            color: #333;
        }
        .table .total {
            font-weight: bold;
            color: #4CAF50;
        }
        .footer {
            text-align: center;
            font-size: 12px;
            color: #777;
            margin-top: 20px;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>Monthly Payslip</h1>
            <p>For the period: January 2025</p>
        </div>
        
        <!-- Employee Information using Flexbox -->
        <div class="employee-info">
            <div>
                <h3>Employee Information</h3>
                <div><strong>Name:</strong> John Doe</div>
                <div><strong>Employee ID:</strong> 123456</div>
                <div><strong>Designation:</strong> Software Engineer</div>
            </div>
            <div>
                <h3>Company Information</h3>
                <div><strong>Company:</strong> XYZ Corp</div>
                <div><strong>Department:</strong> IT</div>
                <div><strong>Location:</strong> New York, USA</div>
            </div>
        </div>

        <!-- Earnings Section using CSS Grid -->
        <div class="section">
            <h3>Earnings</h3>
            <div class="table">
                <div class="header">Basic Salary</div>
                <div>$3,500.00</div>
                <div class="header">Bonus</div>
                <div>$500.00</div>
                <div class="header">Overtime</div>
                <div>$150.00</div>
                <div class="header total">Total Earnings</div>
                <div class="total">$4,150.00</div>
            </div>
        </div>

        <!-- Deductions Section using CSS Grid -->
        <div class="section">
            <h3>Deductions</h3>
            <div class="table">
                <div class="header">Tax</div>
                <div>$450.00</div>
                <div class="header">Insurance</div>
                <div>$100.00</div>
                <div class="header total">Total Deductions</div>
                <div class="total">$550.00</div>
            </div>
        </div>

        <!-- Net Salary Section using CSS Grid -->
        <div class="section">
            <h3>Net Salary</h3>
            <div class="table">
                <div class="header total">Net Salary</div>
                <div class="total">$3,600.00</div>
            </div>
        </div>

        <div class="footer">
            <p>&copy; 2025 Company Name. All rights reserved.</p>
        </div>
    </div>
</body>
</html>
`

	page := wkhtmltopdf.NewPageReader(strings.NewReader(htmlContent))
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("output.pdf")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("PDF Created Successfully")
}
