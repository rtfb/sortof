// var _ = Underscore.load();

function onOpen() {
  var ui = SpreadsheetApp.getUi();
  // Or DocumentApp or FormApp.
  ui.createMenu('Spam Everyone')
      .addItem('Send emails', 'sendEmails')
      .addToUi();
}

function getMetadata() {
  var date = SpreadsheetApp.getActiveSpreadsheet().getSheetByName("metadata").getRange("B2").getValue();
  var signature = SpreadsheetApp.getActiveSpreadsheet().getSheetByName("metadata").getRange("B3").getValue();
  return {
    date: date,
    signature: signature
  }
}

function getData(sheetName) {
  meta = getMetadata()
  var allData = [];
  var sheet = SpreadsheetApp.getActiveSheet();
  var selection = sheet.getSelection();
  var ranges =  selection.getActiveRangeList().getRanges();
  for (var i = 0; i < ranges.length; i++) {
    var dataRange = sheet.getRange(ranges[i].getRowIndex(), 1, ranges[i].getHeight(), 14);
    var rangeValues = dataRange.getValues();
    for (var j = 0; j < rangeValues.length; j++) {
      //Logger.log('name1: ' + rangeValues[j][0]);
      allData.push(rangeValues[j]);
    }
  }
  var records = [];
  for (var i = 0; i < allData.length; i++) {
    var t = allData[i];
    // Logger.log('name: ' + t[0]);
    var name = t[0];
    var email = t[1];
    var filename = t[2];
    records.push({
      recipient: email,
      name: name,
      filename: filename
    });
  }
  return {
    records: records,
    meta: meta
  }
}

function createLogSheet() {
  var timeZone = "UTC";
  dt = Utilities.formatDate(new Date(), timeZone, "yyyy-MM-dd | HH:mm:ss");
  var logSheet = SpreadsheetApp.getActiveSpreadsheet().insertSheet("Emails log - " + dt + " (UTC)");
  logSheet.setColumnWidth(1, 300);
  var header = logSheet.getRange("A1:B1");
  header.setHorizontalAlignments([["center", "center"]]);
  header.setFontWeights([["bold", "bold"]]);
  header.setValues([["Message", "Status"]]);
  return logSheet
}

function sendEmails() {
  var data = getData("Sheet1");
  var logs = [];
  for (var i = 0; i < data.records.length; i++) {
    dataItem = data.records[i];
    // Logger.log(dataItem);
    var recipient = "nachema.june@gmail.com";
    // var cc = dataItem.recipient + ",vytas@rtfb.lt";
    var cc = "vytas@rtfb.lt";
    var filename = dataItem.filename;
    var subject = "Test email";
    var body = "Hello " + dataItem.name + "(CC: " + dataItem.recipient + "),\n\n\
This is a test email.\n\n" + data.meta.date + "\n" + data.meta.signature;
    var htmlBody = "Hello " + dataItem.name + "(CC: " + dataItem.recipient + "),<br/>\
<br/>\
This is a test email.<br/><br/>" + data.meta.date + "<br/>" + data.meta.signature;

    try {
      var files = DriveApp.getFilesByName(filename);
      if (!files.hasNext()) {
        logs.push({email: cc, status: "no file " + filename});
        continue;
      }
      file = files.next();
      if (files.hasNext()) {
        logs.push({email: cc, status: "multiple files for " + filename});
        continue;
      }
      // GmailApp.createDraft(recipient, subject, body, {
      GmailApp.sendEmail(recipient, subject, body, {
        attachments: file,
        htmlBody: htmlBody,
        cc: cc
      });
      logs.push({email: cc, status: "OK"});
    } catch(e) {
      logs.push({email: cc, status: "exception: " + e});
      Logger.log(dataItem.name + ": error with email (" + cc + "). " + e);
    }
  }
  var logSheet = createLogSheet();
  var rowNumber = 1; // row numbers are 1-based
  const messageCol = 1;
  const statusCol = 2;
  for (var i = 0; i < logs.length; i++) {
    rowNumber++; // skip the header
    logSheet.getRange(rowNumber, messageCol).setValue("sending email to " + logs[i].email + "...");
    logSheet.getRange(rowNumber, statusCol).setValue(logs[i].status);
  }
}
