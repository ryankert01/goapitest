#!/bin/bash

# Check if Java is installed
if ! command -v java &> /dev/null; then
    echo "Java is not installed. Installing Java..."
    sudo apt-get update
    sudo apt-get install default-jdk
fi

# Check if JMeter is already installed
if ! command -v jmeter &> /dev/null; then
    echo "JMeter is not installed. Installing JMeter..."
    # Download JMeter (replace with the latest version link if needed)
    JMETER_URL="https://dlcdn.apache.org//jmeter/binaries/apache-jmeter-5.6.3.zip"
    JMETER_FILE="apache-jmeter-5.6.3.zip"

    wget $JMETER_URL -O $JMETER_FILE

    # Extract JMeter
    unzip $JMETER_FILE
    rm $JMETER_FILE  # Clean up the zip file
    cp -r apache-jmeter-5.6.3/ $HOME
    rm -rf apache-jmeter-5.6.3/

    # Optional: Set environment variables 
    JMETER_HOME="$HOME/apache-jmeter-5.6.3"  # Adjust if you extracted to a different location
    echo 'export JMETER_HOME='$JMETER_HOME >> ~/.bashrc 
    echo 'export PATH=$PATH:$JMETER_HOME/bin' >> ~/.bashrc
    source ~/.bashrc

    echo "JMeter installation complete!"
fi

echo "Running JMeter tests..."
# make sure log.jtl and report folder are not present
rm -f log.jtl
rm -rf ./test/report
JVM_ARGS="-Xms3g -Xmx3g" jmeter -n -t ./test/test.jmx -l log.jtl -e -o ./test/report
echo "JMeter tests completed!"